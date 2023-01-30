package registry

import (
	"github.com/driver005/gateway/billing/credit"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/invoiceItem"
	"github.com/driver005/gateway/billing/plan"
	"github.com/driver005/gateway/billing/quote"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/billing/usageRecord"
)

func (m *Base) Credit() *credit.Handler {
	if m.credit == nil {
		m.credit = credit.NewHandler(m.r)
	}
	return m.credit
}

func (m *Base) Invoice() *invoice.Handler {
	if m.invoice == nil {
		m.invoice = invoice.NewHandler(m.r)
	}
	return m.invoice
}

func (m *Base) InvoiceItem() *invoiceItem.Handler {
	if m.invoiceItem == nil {
		m.invoiceItem = invoiceItem.NewHandler(m.r)
	}
	return m.invoiceItem
}

func (m *Base) Plan() *plan.Handler {
	if m.plan == nil {
		m.plan = plan.NewHandler(m.r)
	}
	return m.plan
}

func (m *Base) Quote() *quote.Handler {
	if m.quote == nil {
		m.quote = quote.NewHandler(m.r)
	}
	return m.quote
}

func (m *Base) Subscription() *subscription.Handler {
	if m.subscription == nil {
		m.subscription = subscription.NewHandler(m.r)
	}
	return m.subscription
}

func (m *Base) SubscriptionItem() *subscriptionItem.Handler {
	if m.subscriptionItem == nil {
		m.subscriptionItem = subscriptionItem.NewHandler(m.r)
	}
	return m.subscriptionItem
}

func (m *Base) SubscriptionSchedule() *subscriptionSchedule.Handler {
	if m.subscriptionSchedule == nil {
		m.subscriptionSchedule = subscriptionSchedule.NewHandler(m.r)
	}
	return m.subscriptionSchedule
}

func (m *Base) UsageRecord() *usageRecord.Handler {
	if m.usageRecord == nil {
		m.usageRecord = usageRecord.NewHandler(m.r)
	}
	return m.usageRecord
}
