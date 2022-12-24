package service

import "github.com/driver005/gateway/models"

func isOrder(model interface{}) bool {
	switch model.(type) {
	case models.Order:
		return true
	default:
		return false
	}
}

type ShippingMethodTotals struct {
	Price            int                            `json:"price"`
	TaxTotal         int                            `json:"tax_total"`
	Total            int                            `json:"total"`
	Subtotal         int                            `json:"subtotal"`
	OriginalTotal    int                            `json:"original_total"`
	OriginalTaxTotal int                            `json:"original_tax_total"`
	TaxLines         []models.ShippingMethodTaxLine `json:"tax_lines"`
}

type GiftCardAllocation struct {
	Amount     int `json:"amount"`
	UnitAmount int `json:"unit_amount"`
}

type DiscountAllocation struct {
	Amount     int `json:"amount"`
	UnitAmount int `json:"unit_amount"`
}

type LineAllocationsMap struct {
	GiftCard GiftCardAllocation `json:"gift_card"`
	Discount DiscountAllocation `json:"discount"`
}

type TaxCalculationContext struct {
	ShippingAddress *models.Address         `json:"shipping_address"`
	Customer        models.Customer         `json:"customer"`
	Region          models.Region           `json:"region"`
	IsReturn        bool                    `json:"is_return"`
	ShippingMethods []models.ShippingMethod `json:"shipping_methods"`
	AllocationMap   LineAllocationsMap      `json:"allocation_map"`
}

type GetShippingMethodTotalsOptions struct {
	IncludeTax         bool `json:"include_tax"`
	UseTaxLines        bool `json:"use_tax_lines"`
	CalculationContext bool `json:"calculation_context"`
}
