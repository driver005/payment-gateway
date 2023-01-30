package promotion

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/products/coupon"
	"github.com/google/uuid"
)

type PromotionCode struct {
	core.Model

	Active                bool           `json:"active"`
	Code                  string         `json:"code"`
	Coupon                *coupon.Coupon `json:"coupon"`
	ExpiresAt             int            `json:"expires_at,omitempty"`
	MaxRedemptions        int            `json:"max_redemptions,omitempty"`
	FirstTimeTransaction  bool           `json:"first_time_transaction"`
	MinimumAmount         int            `json:"minimum_amount,omitempty"`
	MinimumAmountCurrency string         `json:"minimum_amount_currency,omitempty"`
	TimesRedeemed         int            `json:"times_redeemed"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
