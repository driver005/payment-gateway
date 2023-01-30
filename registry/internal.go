package registry

import (
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/dispute"
	"github.com/driver005/gateway/internal/file"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/internal/payout"
	"github.com/driver005/gateway/internal/refund"
	"github.com/driver005/gateway/internal/setup/attempt"
	setup_intent "github.com/driver005/gateway/internal/setup/intent"
)

func (m *Base) Balance() *balance.Handler {
	if m.balance == nil {
		m.balance = balance.NewHandler(m.r)
	}
	return m.balance
}

func (m *Base) Charge() *charge.Handler {
	if m.charge == nil {
		m.charge = charge.NewHandler(m.r)
	}
	return m.charge
}

func (m *Base) Customer() *customer.Handler {
	if m.customer == nil {
		m.customer = customer.NewHandler(m.r)
	}
	return m.customer
}

func (m *Base) Dispute() *dispute.Handler {
	if m.dispute == nil {
		m.dispute = dispute.NewHandler(m.r)
	}
	return m.dispute
}

func (m *Base) File() *file.Handler {
	if m.file == nil {
		m.file = file.NewHandler(m.r)
	}
	return m.file
}

func (m *Base) PaymentIntent() *intent.Handler {
	if m.intent == nil {
		m.intent = intent.NewHandler(m.r)
	}
	return m.intent
}

func (m *Base) Mandate() *mandate.Handler {
	if m.mandate == nil {
		m.mandate = mandate.NewHandler(m.r)
	}
	return m.mandate
}

func (m *Base) Payout() *payout.Handler {
	if m.payout == nil {
		m.payout = payout.NewHandler(m.r)
	}
	return m.payout
}

func (m *Base) Refund() *refund.Handler {
	if m.refund == nil {
		m.refund = refund.NewHandler(m.r)
	}
	return m.refund
}

func (m *Base) SetupAttempt() *attempt.Handler {
	if m.setup_attempt == nil {
		m.setup_attempt = attempt.NewHandler(m.r)
	}
	return m.setup_attempt
}

func (m *Base) SetupIntent() *setup_intent.Handler {
	if m.setup_intent == nil {
		m.setup_intent = setup_intent.NewHandler(m.r)
	}
	return m.setup_intent
}
