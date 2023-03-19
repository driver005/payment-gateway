package invoiceItem

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/products/price"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Customer() *customer.Handler
	Invoice() *invoice.Handler
	SubscriptionItem() *subscriptionItem.Handler
	Subscription() *subscription.Handler
	Price() *price.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
