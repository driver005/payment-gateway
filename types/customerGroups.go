package types

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type FilterableCustomerGroupProps struct {
	core.Filter
	Q                   string        `json:"q,omitempty"`
	Name                []string      `json:"name,omitempty"`
	DiscountConditionId uuid.NullUUID `json:"discount_condition_id,omitempty"`
}
