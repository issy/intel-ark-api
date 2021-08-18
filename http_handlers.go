package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

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

// Handles requests to the /search endpoint
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
