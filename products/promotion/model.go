package promotion

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/products/coupon"
	"github.com/google/uuid"
)

type PromotionCodeCurrencyOption struct {
	core.Model

	MinimumAmount int32 `json:"minimum_amount"`
}

type PromotionCodesResourceRestrictions struct {
	core.Model

	CurrencyOptions       *PromotionCodeCurrencyOption `json:"currency_options,omitempty"`
	FirstTimeTransaction  bool                         `json:"first_time_transaction"`
	MinimumAmount         int                          `json:"minimum_amount,omitempty"`
	MinimumAmountCurrency string                       `json:"minimum_amount_currency,omitempty"`
}

type PromotionCode struct {
	core.Model

	Active         bool                                `json:"active"`
	Code           string                              `json:"code"`
	Coupon         *coupon.Coupon                      `json:"coupon"`
	ExpiresAt      int                                 `json:"expires_at,omitempty"`
	MaxRedemptions int                                 `json:"max_redemptions,omitempty"`
	Restrictions   *PromotionCodesResourceRestrictions `json:"restrictions"`
	TimesRedeemed  int32                               `json:"times_redeemed"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
