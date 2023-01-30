package plaid

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/payment/bank"
	"github.com/plaid/plaid-go/plaid"
)

type Handler struct {
	r      Registry
	client plaid.APIClient
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Context() *database.DB
	Bank() *bank.Handler
}

func NewHandler(r Registry) *Handler {
	h := Handler{
		r:      r,
		client: *InitPlaid(),
	}
	return &h
}
