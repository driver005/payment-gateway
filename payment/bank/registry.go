package bank

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/utils"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Migrate() bool
	Utils() *utils.Handler
	Customer() *customer.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{r: r}
	if h.r.Migrate() {
		h.Migrate()
	}
	return &h
}
