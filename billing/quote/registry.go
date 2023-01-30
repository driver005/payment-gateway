package quote

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/internal/customer"
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
	Subscription() *subscription.Handler
	SubscriptionSchedule() *subscriptionSchedule.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
