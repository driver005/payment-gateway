package driver

import (
	"github.com/driver005/gateway/driver/paylater/affirm"
	"github.com/driver005/gateway/driver/paylater/klarna"
	"github.com/driver005/gateway/driver/paylater/riverty"
)

type Handler struct {
	r       Registry
	affirm  affirm.Affirm
	riverty riverty.Riverty
	klarna  klarna.CheckoutSrv
}

type Registry interface {
}

func NewHandler(r Registry) *Handler {
	return &Handler{
		r:       r,
		affirm:  InitAffirm(),
		riverty: InitRiverty(),
		klarna:  InitKlarna(),
	}
}
