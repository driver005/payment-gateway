package invoice

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/payment/method"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Charge() *charge.Handler
	Customer() *customer.Handler
	PaymentIntent() *intent.Handler
	PaymentMethod() *method.Handler
	Subscription() *subscription.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
