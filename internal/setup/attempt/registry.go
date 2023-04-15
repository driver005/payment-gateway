package attempt

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/customer"
	setupIntent "github.com/driver005/gateway/internal/setup/intent"
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
	SetupIntent() *setupIntent.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
