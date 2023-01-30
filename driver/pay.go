package driver

import (
	"context"
	"errors"

	"github.com/driver005/gateway/internal/intent"
)

func (h *Handler) Pay(ctx context.Context, m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	var resp *intent.PaymentIntent
	var err error

	if m.PaymentMethod.Type == "acss_debit" {
		return nil, nil
	} else if m.PaymentMethod.Type == "affirm" {
		resp, err = h.Affirm(m)
	} else if m.PaymentMethod.Type == "afterpay_clearpay" {
		resp, err = h.Riverty(m)
	} else if m.PaymentMethod.Type == "alipay" {
		resp, err = h.Alipay(m)
	} else if m.PaymentMethod.Type == "au_becs_debit" {
		return nil, nil
	} else if m.PaymentMethod.Type == "bacs_debit" {
		return nil, nil
	} else if m.PaymentMethod.Type == "bancontact" {
		return nil, nil
	} else if m.PaymentMethod.Type == "blik" {
		return nil, nil
	} else if m.PaymentMethod.Type == "boleto" {
		resp = h.Boleto(m)
	} else if m.PaymentMethod.Type == "btcpay" {
		resp, err = h.BtcPay(m)
	} else if m.PaymentMethod.Type == "card" {
		resp, err = h.Card(m)
	} else if m.PaymentMethod.Type == "card_present" {
		resp, err = h.Card(m)
	} else if m.PaymentMethod.Type == "customer_balance" {
		return nil, nil
	} else if m.PaymentMethod.Type == "eps" {
		resp, err = h.Eps(m)
	} else if m.PaymentMethod.Type == "fpx" {
		resp, err = h.Fpx(m)
	} else if m.PaymentMethod.Type == "giropay" {
		resp, err = h.Giropay(m)
	} else if m.PaymentMethod.Type == "grabpay" {
		_, err = h.GrabPay(m)
		return nil, nil
	} else if m.PaymentMethod.Type == "ideal" {
		resp, err = h.Ideal(m)
	} else if m.PaymentMethod.Type == "interac_present" {
		return nil, nil
	} else if m.PaymentMethod.Type == "klarna" {
		resp, err = h.Klarna(m)
	} else if m.PaymentMethod.Type == "konbini" {
		return nil, nil
	} else if m.PaymentMethod.Type == "link" {
		return nil, nil
	} else if m.PaymentMethod.Type == "oxxo" {
		resp, err = h.Oxxo(m)
	} else if m.PaymentMethod.Type == "p24" {
		return nil, nil
	} else if m.PaymentMethod.Type == "paynow" {
		resp, err = h.PayNow(m)
	} else if m.PaymentMethod.Type == "pix" {
		resp, err = h.Pix(m)
	} else if m.PaymentMethod.Type == "promptpay" {
		resp, err = h.PromptPay(m)
	} else if m.PaymentMethod.Type == "sepa_debit" {
		return nil, nil
	} else if m.PaymentMethod.Type == "sofort" {
		return nil, nil
	} else if m.PaymentMethod.Type == "us_bank_account" {
		return nil, nil
	} else if m.PaymentMethod.Type == "wechat_pay" {
		_, err = h.WeChat(m)
		return nil, nil
	} else {
		return nil, errors.New("payment method not found")
	}
	if err != nil {
		return nil, err
	}

	r, err := h.r.PaymentIntent().Upsert(ctx, resp)
	if err != nil {
		return nil, err
	}

	return r, nil
}
