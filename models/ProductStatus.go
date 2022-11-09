package models

import "database/sql/driver"

// The status of the product
type ProductStatus string

const (
	ProductStatusDraft     ProductStatus = "draft"
	ProductStatusProposed  ProductStatus = "proposed"
	ProductStatusPublished ProductStatus = "published"
	ProductStatusRejected  ProductStatus = "rejected"
)

func (ps *ProductStatus) Scan(value interface{}) error {
	*ps = ProductStatus(value.([]byte))
	return nil
}

func (ps ProductStatus) Value() (driver.Value, error) {
	return string(ps), nil
}
