package intent

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/payment/method"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Customer() *customer.Handler
	PaymentMethod() *method.Handler
	Mandate() *mandate.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
