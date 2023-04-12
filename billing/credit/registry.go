package credit

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/refund"
)

type Handler struct {
	r     Registry
	model interface{}
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Customer() *customer.Handler
	Invoice() *invoice.Handler
	Balance() *balance.Handler
	Refund() *refund.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
