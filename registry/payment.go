package registry

import (
	"github.com/driver005/gateway/payment/bank"
	"github.com/driver005/gateway/payment/card"
	"github.com/driver005/gateway/payment/cash"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/source"
)

func (m *Base) PaymentMethod() *method.Handler {
	if m.method == nil {
		m.method = method.NewHandler(m.r)
	}
	return m.method
}

func (m *Base) Bank() *bank.Handler {
	if m.bank == nil {
		m.bank = bank.NewHandler(m.r)
	}
	return m.bank
}

func (m *Base) Card() *card.Handler {
	if m.card == nil {
		m.card = card.NewHandler(m.r)
	}
	return m.card
}

func (m *Base) Cash() *cash.Handler {
	if m.cash == nil {
		m.cash = cash.NewHandler(m.r)
	}
	return m.cash
}

func (m *Base) Source() *source.Handler {
	if m.source == nil {
		m.source = source.NewHandler(m.r)
	}
	return m.source
}
