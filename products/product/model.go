package product

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/tax"
	"github.com/lib/pq"
)

type PackageDimensions struct {
	core.Model

	Height float32 `json:"height,omitempty"`
	Length float32 `json:"length,omitempty"`
	Weight float32 `json:"weight,omitempty"`
	Width  float32 `json:"width,omitempty"`
}

type Product struct {
	core.Model

	Active              bool               `json:"active,omitempty"`
	Description         string             `json:"description,omitempty"`
	Images              pq.StringArray     `json:"images,omitempty" database:"type:varchar(64)[]"`
	Name                string             `json:"name,omitempty"`
	PackageDimensions   *PackageDimensions `json:"package_dimensions,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Shippable           bool               `json:"shippable,omitempty"`
	StatementDescriptor string             `json:"statement_descriptor,omitempty"`
	TaxCode             *tax.TaxCode       `json:"tax_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UnitLabel           string             `json:"unit_label,omitempty"`
	Updated             int                `json:"updated,omitempty"`
	Url                 string             `json:"url,omitempty"`
}
