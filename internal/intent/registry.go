package intent

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Pay(ctx context.Context, i *PaymentIntent) (*PaymentIntent, error)
	Migrate() bool
	PaymentMethod() *method.Handler
	Customer() *customer.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
