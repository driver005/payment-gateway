package discount

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/products/coupon"
	"github.com/driver005/gateway/products/promotion"
	"github.com/google/uuid"
)

type Discount struct {
	core.Model

	CheckoutSession string         `json:"checkout_session,omitempty"`
	Coupon          *coupon.Coupon `json:"coupon"`
	End             int            `json:"end,omitempty"`
	Invoice         string         `json:"invoice,omitempty"`
	InvoiceItem     string         `json:"invoice_item,omitempty"`
	Start           int32          `json:"start"`
	Subscription    string         `json:"subscription,omitempty"`

	CustomerId      uuid.UUID                `json:"customer_id" database:"default:null"`
	Customer        *customer.Customer       `json:"customer,omitempty" database:"foreignKey:customer_id"`
	PromotionCodeId uuid.UUID                `json:"promotion_code_id" database:"default:null"`
	PromotionCode   *promotion.PromotionCode `json:"promotion_code,omitempty" database:"foreignKey:promotion_code_id"`
}
