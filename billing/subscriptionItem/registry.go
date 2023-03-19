package subscriptionItem

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/products/price"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
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
