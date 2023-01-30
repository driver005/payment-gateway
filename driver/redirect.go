package driver

import (
	"github.com/driver005/gateway/driver/redirect/bancontact"
	"github.com/driver005/gateway/driver/redirect/blik"
	"github.com/driver005/gateway/driver/redirect/eps"
	"github.com/driver005/gateway/driver/redirect/fpx"
	"github.com/driver005/gateway/driver/redirect/giropay"
	"github.com/driver005/gateway/driver/redirect/ideal"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/spf13/viper"
)

//Bancontact

func InitBancontact() *bancontact.Bankcontact {
	return bancontact.NewBankcontact(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Bancontact(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.bancontact.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}

//Blik

func InitBlik() *blik.Blik {
	return blik.NewBlik(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Blik(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.blik.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}

//Eps

func InitEps() *eps.Eps {
	return eps.NewEps(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Eps(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.eps.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}

//Fpx

func InitFpx() *fpx.Fpx {
	return fpx.NewFpx(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Fpx(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.fpx.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}

//GiroPay

func InitGiroPay() *giropay.GiroPay {
	return giropay.NewGiroPay(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Giropay(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.giropay.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}

//Ideal

func InitIdeal() *ideal.Ideal {
	return ideal.NewIdeal(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Ideal(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.ideal.Payment(m)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
			Url: payload.NextAction.RedirectToURL.URL,
		},
	}

	return m, nil
}
