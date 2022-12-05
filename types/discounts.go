package types

import (
	"github.com/driver005/gateway/core"
)

type FilterableDiscountProps struct {
	core.Filter
	Q          string      `json:"q,omitempty"`
	IsDynamic  bool        `json:"is_dynamic,omitempty"`
	IsDisabled bool        `json:"is_disabled,omitempty"`
	Rule       interface{} `json:"rule,omitempty"`
}
