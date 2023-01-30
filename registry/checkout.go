package registry

import "github.com/driver005/gateway/checkout"

func (m *Base) Checkout() *checkout.Handler {
	if m.checkout == nil {
		m.checkout = checkout.NewHandler(m.r)
	}
	return m.checkout
}
