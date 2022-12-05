package handler

import (
	"context"

	"github.com/driver005/database"
	db "github.com/driver005/gateway/database"
)

type Handler struct {
	r Registry
	n Nbxplorer
	b BtcPay
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	ClientManager() *db.Handler
}

func NewHandler(r Registry, host string) *Handler {
	return &Handler{
		r: r,
		n: *NewNbxplorer(r, host),
		b: *NewBtcPay(r),
	}
}
