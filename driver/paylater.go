package driver

import (
	"net/url"
	"time"

	"github.com/driver005/gateway/driver/paylater/affirm"
	"github.com/driver005/gateway/driver/paylater/klarna"
	"github.com/driver005/gateway/driver/paylater/riverty"
)

//
// Affirm
//

func InitAffirm() affirm.Affirm {
	uri, _ := url.Parse(affirm.API)
	conf := affirm.Config{
		BaseURL: uri,
		APIKey: "",
		Timeout: time.Second * 10,
	}

	client := affirm.NewClient(conf)

	return affirm.NewAffirm(client)
}

func (h *Handler) Affirm(m *affirm.DirectCheckout) (string, error) {
	resp, err := h.affirm.CreateDirectCheckout(m)
	if err != nil {
		return "", err
	}

	return resp.RedirectUrl, nil
}

//
// Riverty
//

func InitRiverty() riverty.Riverty {
	uri, _ := url.Parse(riverty.API)
	conf := riverty.Config{
		BaseURL: uri,
		APIKey: "",
		Timeout: time.Second * 10,
	}

	client := riverty.NewClient(conf)

	return riverty.NewRiverty(client)
}

func (h *Handler) Riverty(m *riverty.Checkout) (string, error) {
	resp, err := h.riverty.CreateCheckout(m)
	if err != nil {
		return "", err
	}

	return resp.SecureLoginURL, nil
}

//
// Klarna
//

func InitKlarna() klarna.CheckoutSrv {
	uri, _ := url.Parse(riverty.API)
	conf := klarna.Config{
		BaseURL: uri,
		APIUsername: "",
		APIPassword: "",
		Timeout: time.Second * 10,
	}

	client := klarna.NewClient(conf)

	return klarna.NewCheckoutSrv(client)
}

func (h *Handler) Klarna(m *klarna.CheckoutOrder) (string, error) {
	resp, err := h.klarna.CreateNewOrder(m)
	if err != nil {
		return "", err
	}

	return resp.HtmlSnippet, nil
}
