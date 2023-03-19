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

	Amount          int    `json:"amount,omitempty"`
	CheckoutSession string `json:"checkout_session,omitempty"`
	End             int    `json:"end,omitempty"`
	Invoice         string `json:"invoice,omitempty"`
	InvoiceItem     string `json:"invoice_item,omitempty"`
	Start           int    `json:"start,omitempty"`
	Subscription    string `json:"subscription,omitempty"`

	CustomerId      *uuid.UUID               `json:"customer_id,omitempty" swaggerignore:"true"`
	Customer        *customer.Customer       `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
	CouponId        *uuid.UUID               `json:"coupon_id,omitempty" swaggerignore:"true"`
	Coupon          *coupon.Coupon           `json:"coupon,omitempty" database:"foreignKey:coupon_id" swaggertype:"primitive,string" format:"uuid"`
	PromotionCodeId *uuid.UUID               `json:"promotion_code_id,omitempty" swaggerignore:"true"`
	PromotionCode   *promotion.PromotionCode `json:"promotion_code,omitempty" database:"foreignKey:promotion_code_id" swaggertype:"primitive,string" format:"uuid"`
}
