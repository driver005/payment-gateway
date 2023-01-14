package handler

import (
	"context"

	"github.com/driver005/database"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/service"
)

type Handler struct {
	r Registry
	n *Nbxplorer
	b *BtcPay
	a *ApplePay
	g *GooglePay
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	ClientManager() *db.Handler
	Service() *service.Handler
	Logger() *logger.Logger
	// Applepay() *applepay.Merchant
}

func NewHandler(r Registry, host string) *Handler {
	return &Handler{
		r: r,
		n: NewNbxplorer(r, host),
		b: NewBtcPay(r),
		a: NewApplePay(),
		g: NewGooglePay(),
	}
}
