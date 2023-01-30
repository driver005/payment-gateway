package charge

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Customer() *customer.Handler
	Balance() *balance.Handler
	PaymentMethod() *method.Handler
	Utils() *utils.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
