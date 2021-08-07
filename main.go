package main

import (
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"

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

type response struct {
	TotalRecords int                `json:"totalRecords"`
	FirstRecord  int                `json:"firstRecord"`
	Products     []populatedProduct `json:"products"`
}

func getProductIDs(searchTerm string) []rawProduct {
	if len(searchTerm) == 0 {
		return make([]rawProduct, 0)
	}

	var products []rawProduct
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

func convertIDToValue(resourceID int) string {
	var result string
	query := fmt.Sprintf("SELECT Value FROM Resource WHERE fkResourceID = %d", resourceID)
	db.Raw(query).Scan(&result)
	return result
}

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

func intFromParams(params *url.Values, key string, defaultValue int) int {
	if val := params.Get(key); val != "" {
		if numericVal, err := strconv.Atoi(val); err == nil {
			return numericVal
		}
	}
	return defaultValue
}

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

func searchHandler(c *gin.Context) {
	searchTerm := c.Param("search-term")
	queryParams := c.Request.URL.Query()
	firstRecord := constrainInt(intFromParams(&queryParams, "firstRecord", 0), 0, math.MaxInt64)
	count := constrainInt(intFromParams(&queryParams, "count", 50), 1, 50)

	rawProducts := getProductIDs(searchTerm)
	populatedProducts := make([]populatedProduct, 0)
	for i := firstRecord; i < len(rawProducts) && i < firstRecord+count; i++ {
		productMetadata := rawProducts[i]
		fullProduct := getProductSpecs(productMetadata)
		populatedProducts = append(populatedProducts, fullProduct)
	}

	c.JSON(
		http.StatusOK,
		response{
			TotalRecords: len(rawProducts),
			FirstRecord:  firstRecord,
			Products:     populatedProducts,
		},
	)
}

func main() {
	db = initDB()
	router := gin.Default()
	router.GET("/search/:search-term", searchHandler)
	router.Run("localhost:8000")
}
