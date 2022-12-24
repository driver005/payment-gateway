package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableProductVariantProps struct {
	core.Filter
	Q                 string          `json:"q,omitempty"`
	Title             []string        `json:"title,omitempty"`
	ProductId         []uuid.NullUUID `json:"product_id,omitempty"`
	Sku               []string        `json:"sku,omitempty"`
	Barcode           []string        `json:"barcode,omitempty"`
	Ean               []string        `json:"ean,omitempty"`
	Upc               []string        `json:"upc,omitempty"`
	InventoryQuantity core.Number     `json:"inventory_quantity,omitempty"`
	AllowBackorder    bool            `json:"allow_backorder,omitempty"`
	ManageInventory   bool            `json:"manage_inventory,omitempty"`
	HsCode            []string        `json:"hs_code,omitempty"`
	OriginCountry     []string        `json:"origin_country,omitempty"`
	MidCode           []string        `json:"mid_code,omitempty"`
	Material          string          `json:"material,omitempty"`
	Weight            core.Number     `json:"weight,omitempty"`
	Length            core.Number     `json:"length,omitempty"`
	Height            core.Number     `json:"height,omitempty"`
	Width             core.Number     `json:"width,omitempty"`
}
