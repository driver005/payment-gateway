package coupon

import "github.com/driver005/gateway/core"

type CouponAppliesTo struct {
	core.Model

	Products []string `json:"products,omitempty"`
}

type CouponCurrencyOption struct {
	core.Model

	AmountOff int32 `json:"amount_off,omitempty"`
}

type Coupon struct {
	core.Model

	AmountOff        int                   `json:"amount_off,omitempty"`
	AppliesTo        *CouponAppliesTo      `json:"applies_to,omitempty"`
	Currency         string                `json:"currency,omitempty"`
	CurrencyOptions  *CouponCurrencyOption `json:"currency_options,omitempty"`
	Duration         string                `json:"duration,omitempty"`
	DurationInMonths int                   `json:"duration_in_months,omitempty"`
	MaxRedemptions   int                   `json:"max_redemptions,omitempty"`
	Name             string                `json:"name,omitempty"`
	PercentOff       float64               `json:"percent_off,omitempty"`
	RedeemBy         int                   `json:"redeem_by,omitempty"`
	TimesRedeemed    int32                 `json:"times_redeemed"`
	Valid            bool                  `json:"valid,omitempty"`
}
