package product

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

type PackageDimensions struct {
	core.Model

	Height float32 `json:"height"`
	Length float32 `json:"length"`
	Weight float32 `json:"weight"`
	Width  float32 `json:"width"`
}

type Product struct {
	core.Model

	Active              bool              `json:"active"`
	Description         string            `json:"description,omitempty"`
	Images              []string          `json:"images"`
	Name                string            `json:"name"`
	PackageDimensions   PackageDimensions `json:"package_dimensions,omitempty"`
	Shippable           bool              `json:"shippable,omitempty"`
	StatementDescriptor string            `json:"statement_descriptor,omitempty"`
	TaxCode             *tax.TaxCode      `json:"tax_code,omitempty"`
	UnitLabel           string            `json:"unit_label,omitempty"`
	Updated             int               `json:"updated"`
	Url                 string            `json:"url,omitempty"`

	DefaultPriceId uuid.UUID    `json:"default_price_id,omitempty"`
	DefaultPrice   *price.Price `json:"default_price,omitempty"  database:"foreignKey:default_price_id"`
}
