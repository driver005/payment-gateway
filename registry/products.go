package registry

import (
	"github.com/driver005/gateway/products/coupon"
	"github.com/driver005/gateway/products/discount"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/products/product"
	"github.com/driver005/gateway/products/promotion"
)

func (m *Base) Coupon() *coupon.Handler {
	if m.coupon == nil {
		m.coupon = coupon.NewHandler(m.r)
	}
	return m.coupon
}

func (m *Base) Discount() *discount.Handler {
	if m.discount == nil {
		m.discount = discount.NewHandler(m.r)
	}
	return m.discount
}

func (m *Base) Price() *price.Handler {
	if m.price == nil {
		m.price = price.NewHandler(m.r)
	}
	return m.price
}

func (m *Base) Product() *product.Handler {
	if m.product == nil {
		m.product = product.NewHandler(m.r)
	}
	return m.product
}

func (m *Base) Promotion() *promotion.Handler {
	if m.promotion == nil {
		m.promotion = promotion.NewHandler(m.r)
	}
	return m.promotion
}
