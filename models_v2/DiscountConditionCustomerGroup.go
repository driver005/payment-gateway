package models

import "github.com/google/uuid"

// DiscountConditionCustomerGroup - Associates a discount condition with a customer group
type DiscountConditionCustomerGroup struct {
	// The ID of the Product Tag
	CustomerGroupId uuid.NullUUID `json:"customer_group_id"`

	CustomerGroup *CustomerGroup `json:"customer_group" database:"foreignKey:id;references:customer_group_id"`

	// The ID of the Discount Condition
	ConditionId uuid.NullUUID `json:"condition_id"`

	DiscountCondition *DiscountCondition `json:"discount_condition" database:"foreignKey:id;references:condition_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
