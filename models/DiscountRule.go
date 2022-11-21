package models

import "github.com/driver005/gateway/core"

// DiscountRule - Holds the rules that governs how a Discount is calculated when applied to a Cart.
type DiscountRule struct {
	core.Model

	// The type of the Discount, can be `fixed` for discounts that reduce the price by a fixed amount, `percentage` for percentage reductions or `free_shipping` for shipping vouchers.
	Type string `json:"type"`

	// A short description of the discount
	Description string `json:"description" database:"default:null"`

	// The value that the discount represents; this will depend on the type of the discount
	Value int32 `json:"value"`

	// The scope that the discount should apply to.
	Allocation string `json:"allocation" database:"default:null"`

	// A set of conditions that can be used to limit when  the discount can be used. Available if the relation `conditions` is expanded.
	Conditions []DiscountCondition `json:"conditions" database:"foreignKey:id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
