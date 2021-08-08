package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

var db *gorm.DB

type rawSpecField struct {
	SectionID int
	NameID    int
	ValueID   int
}

type rawProduct struct {
	ArkID int
	ID    int
	Title string
}

type populatedProduct struct {
	ArkID int        `json:"id"`
	Title string     `json:"name"`
	Specs [][]string `json:"specs"`
}

// Fetches all the known product IDs and names from the resource DB that matches the search term
func getProductIDs(searchTerm string) []rawProduct {
	if len(searchTerm) == 0 {
		return make([]rawProduct, 0)
	}

	// var products []rawProduct
	products := make([]rawProduct, 0)
	query := `SELECT
  Product.ARKID AS ArkID,
  Product.pkProductID AS ID,
  Resource.Value AS Title
  FROM Resource
  LEFT JOIN Product
  ON Product.fkProductNameID = Resource.fkResourceID
  WHERE Resource.Value LIKE "%` + searchTerm + `%"
  AND Resource.Value NOT LIKE "%Tray%"
  AND Resource.Value NOT LIKE "%China%"
  AND Resource.VALUE NOT LIKE "%Boxed%"
  AND Product.pkProductID IS NOT NULL;
  `
	db.Raw(query).Scan(&products)
	return products
}

// Fetches a string from the resource DB by ID
func convertIDToValue(resourceID int) string {
	var result string
	query := fmt.Sprintf("SELECT Value FROM Resource WHERE fkResourceID = %d", resourceID)
	db.Raw(query).Scan(&result)
	return result
}

// Converts a rawProduct to a populatedProduct
// Fetches all the spec rows for the product
func getProductSpecs(product rawProduct) populatedProduct {
	var rawSpecFields []rawSpecField
	query := fmt.Sprintf(`
  SELECT DISTINCT
  Specification.fkSectionID AS SectionID,
  Specification.fkNameID AS NameID,
  Specification.fkValueID AS ValueID
  FROM Specification
  LEFT JOIN ProductSpecificationLink
  ON ProductSpecificationLink.fkProductID = %d
  WHERE Specification.pkSpecificationID = ProductSpecificationLink.fkSpecificationID;
  `, product.ID)
	db.Raw(query).Scan(&rawSpecFields)

	populatedSpecFields := make([][]string, 0)
	for i := 0; i < len(rawSpecFields); i++ {
		rawSpecRow := rawSpecFields[i]
		specRow := make([]string, 3)
		specRow[0] = convertIDToValue(rawSpecRow.SectionID)
		specRow[1] = convertIDToValue(rawSpecRow.NameID)
		specRow[2] = convertIDToValue(rawSpecRow.ValueID)

		populatedSpecFields = append(
			populatedSpecFields,
			specRow,
		)
	}

	return populatedProduct{
		ArkID: product.ArkID,
		Title: product.Title,
		Specs: populatedSpecFields,
	}
}

// Fetch a single int from URL query params by key
func intFromParams(params *url.Values, key string, defaultValue int) int {
	if val := params.Get(key); val != "" {
		if numericVal, err := strconv.Atoi(val); err == nil {
			return numericVal
		}
	}
	return defaultValue
}

// Ensures an int is between a min and max value
func constrainInt(num int, min int, max int) int {
	if num < min {
		return min
	} else if max < num {
		return max
	} else {
		return num
	}
}

func initDB() *gorm.DB {
	new_db, _ := gorm.Open(sqlite.Open("ark-en.sqlite"), &gorm.Config{})
	return new_db
}

// Returns a link header string for pagination purposes
func getPageLinks(urlTemplate string, currentPage int, totalPages int) string {
	linkTemplate := fmt.Sprintf("<%s>; rel=\"%s\"", urlTemplate, "%s")
	pageUrls := make([]string, 0)
	if totalPages == 1 {
		// Single page - pagination not possible
		return ""
	}
	if currentPage != 1 {
		// Not first page
		// Add link to previous page
		pageUrls = append(pageUrls, fmt.Sprintf(linkTemplate, currentPage-1, "prev"))
	}
	if currentPage != totalPages {
		// Not last page
		// Add link to next page
		pageUrls = append(pageUrls, fmt.Sprintf(linkTemplate, currentPage+1, "next"))
	}
	// Add link to last page
	pageUrls = append(pageUrls, fmt.Sprintf(linkTemplate, totalPages, "last"))
	// Add link to first page
	pageUrls = append(pageUrls, fmt.Sprintf(linkTemplate, 1, "first"))
	return strings.Join(pageUrls[:], ", ")
}

func searchHandler(c *gin.Context) {
	queryParams := c.Request.URL.Query()
	searchTerm := queryParams.Get("query")
	pageSize := 50

	rawProducts := getProductIDs(searchTerm)
	if len(rawProducts) == 0 {
		c.JSON(http.StatusOK, rawProducts)
		return
	}

	totalPages := len(rawProducts) / pageSize
	if (len(rawProducts) % pageSize) > 0 {
		totalPages += 1
	}
	currentPage := constrainInt(intFromParams(&queryParams, "page", 1), 1, totalPages)
	urlTemplate := fmt.Sprintf("%s?query=%s&page=%s", c.FullPath(), searchTerm, "%d")
	pageLinks := getPageLinks(urlTemplate, currentPage, totalPages)
	if pageLinks != "" {
		c.Writer.Header().Set("Link", pageLinks)
	}

	populatedProducts := make([]populatedProduct, 0)
	for i := (currentPage - 1) * pageSize; i < len(rawProducts) && i < (currentPage*pageSize); i++ {
		productMetadata := rawProducts[i]
		fullProduct := getProductSpecs(productMetadata)
		populatedProducts = append(populatedProducts, fullProduct)
	}

	c.JSON(http.StatusOK, populatedProducts)
}

func main() {
	db = initDB()
	router := gin.Default()
	router.GET("/search", searchHandler)
	router.Run("localhost:8000")
}
