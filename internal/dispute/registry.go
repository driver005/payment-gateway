package dispute

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/intent"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	PaymentIntent() *intent.Handler
	Charge() *charge.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
