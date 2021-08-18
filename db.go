package main

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db, _ = gorm.Open(sqlite.Open("db/ark-en.sqlite"), &gorm.Config{})

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
