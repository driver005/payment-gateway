package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// TaxRate - A Tax Rate can be used to associate a certain rate to charge on products within a given Region
type TaxRate struct {
	core.Model

	// The numeric rate to charge
	Rate float64 `json:"rate" database:"default:null"`

	// A code to identify the tax type by
	Code string `json:"code" database:"default:null"`

	// A human friendly name for the tax
	Name string `json:"name"`

	// The id of the Region that the rate belongs to
	RegionId uuid.NullUUID `json:"region_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *Region `json:"region" database:"foreignKey:id;references:region_id"`
}
