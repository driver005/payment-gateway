package coupon

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

type Coupon struct {
	core.Model

	AmountOff        int            `json:"amount_off,omitempty"`
	AppliesTo        pq.StringArray `json:"applies_to,omitempty" database:"type:varchar(64)[]"`
	Currency         string         `json:"currency,omitempty"`
	Duration         Duration       `json:"duration,omitempty"`
	DurationInMonths int            `json:"duration_in_months,omitempty"`
	MaxRedemptions   int            `json:"max_redemptions,omitempty"`
	Name             string         `json:"name,omitempty"`
	PercentOff       float64        `json:"percent_off,omitempty"`
	RedeemBy         int            `json:"redeem_by,omitempty"`
	TimesRedeemed    int            `json:"times_redeemed"`
	Valid            bool           `json:"valid,omitempty"`
}
