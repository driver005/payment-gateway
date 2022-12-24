package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// DiscountCondition - Holds rule conditions for when a discount is applicable
type DiscountCondition struct {
	core.Model

	// The type of the Condition
	Type DiscountConditionType `json:"type"`

	// The operator of the Condition
	Operator string `json:"operator"`

	// The ID of the discount rule associated with the condition
	DiscountRuleId uuid.NullUUID `json:"discount_rule_id"`

	DiscountRule *DiscountRule `json:"discount_rule" database:"foreignKey:id;references:discount_rule_id"`

	// products associated with this condition if type = products. Available if the relation `products` is expanded.
	Products []Product `json:"products" database:"foreignKey:id"`

	// product types associated with this condition if type = product_types. Available if the relation `product_types` is expanded.
	ProductTypes []ProductType `json:"product_types" database:"foreignKey:id"`

	// product tags associated with this condition if type = product_tags. Available if the relation `product_tags` is expanded.
	ProductTags []ProductTag `json:"product_tags" database:"foreignKey:id"`

	// product collections associated with this condition if type = product_collections. Available if the relation `product_collections` is expanded.
	ProductCollections []ProductCollection `json:"product_collections" database:"foreignKey:id"`

	// customer groups associated with this condition if type = customer_groups. Available if the relation `customer_groups` is expanded.
	CustomerGroups []CustomerGroup `json:"customer_groups" database:"foreignKey:id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
