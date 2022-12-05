package types

import (
	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

type FilterableProductTaxRateProps struct {
	core.Filter
	ProductId []uuid.NullUUID `json:"product_id,omitempty"`
	RateId    []uuid.NullUUID `json:"rate_id,omitempty"`
}

type FilterableShippingTaxRateProps struct {
	core.Filter
	ShippingOptionId []uuid.NullUUID `json:"shipping_option_id,omitempty"`
	RateId           []uuid.NullUUID `json:"rate_id,omitempty"`
}

type FilterableTaxRateProps struct {
	core.Filter
	RegionId []uuid.NullUUID `json:"region_id,omitempty"`
	RateId   []uuid.NullUUID `json:"rate_id,omitempty"`
	Code     core.String     `json:"code,omitempty"`
	Name     []string        `json:"name,omitempty"`
	Rate     core.Number     `json:"rate,omitempty"`
}
