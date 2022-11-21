package models

import "github.com/gofrs/uuid"

// DiscountConditionProductTag - Associates a discount condition with a product tag
type DiscountConditionProductTag struct {
	// The ID of the Product Tag
	ProductTagId uuid.NullUUID `json:"product_tag_id"`

	ProductTag *ProductTag `json:"product_tag" database:"foreignKey:id;references:product_tag_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" database:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
