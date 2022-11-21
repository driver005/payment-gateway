package models

import "github.com/gofrs/uuid"

// DiscountConditionProductCollection - Associates a discount condition with a product collection
type DiscountConditionProductCollection struct {
	// The ID of the Product Collection
	ProductCollectionId uuid.NullUUID `json:"product_collection_id"`

	ProductCollection *ProductCollection `json:"product_collection" database:"foreignKey:id;references:product_collection_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" database:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
