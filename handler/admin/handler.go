package admin

import (
	"github.com/driver005/gateway/handler"
	"github.com/driver005/gateway/helper"
)

const (
	Currencies = "/currencies"
	BtcpayPath = "/btcpay"
)

type Handler struct {
	c Currency
}

func NewHandler(r handler.Registry) *Handler {
	return &Handler{
		c: Currency{r: r},
	}
}

func (h *Handler) CurrencyRoutes(public *helper.RouteGroup) {
	public.GET(Currencies, h.c.ListCurrency)
	public.GET(Currencies+"/:code", h.c.GetCurrency)
	public.POST(Currencies+"/:code", h.c.UpdateCurrency)
}
