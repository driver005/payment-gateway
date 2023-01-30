package driver

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/driver005/gateway/driver/realtime/paynow"
	"github.com/driver005/gateway/driver/realtime/pix"
	"github.com/driver005/gateway/driver/realtime/promptpay"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
)

func (h *Handler) PayNow(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payee := paynow.NewMobile(m.PaymentMethod.BillingDetails.Phone)

	result := payee.Generate(float32(m.Amount), "ABCDEFG", false, time.Time{})

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "paynow_display_qr_code",
		PaynowDisplayQrCode: &next.PaymentIntentNextActionPaynowDisplayQrCode{
			Data:        result,
			ImageUrlPng: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "paynow", "png"),
			ImageUrlSvg: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "paynow", "svg"),
		},
	}

	return m, nil
}

func (h *Handler) Pix(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	options := pix.Options{
		Name:          m.PaymentMethod.BillingDetails.Name,
		Key:           m.PaymentMethod.BillingDetails.Email,
		City:          m.PaymentMethod.BillingDetails.Address.City,
		Amount:        float64(m.Amount), // optional
		Description:   m.Description,     // optional
		TransactionID: m.Id.String(),     // optional
	}

	result, err := pix.Pix(options)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "pix_display_qr_code",
		PixDisplayQrCode: &next.PaymentIntentNextActionPixDisplayQrCode{
			Data:        result,
			ImageUrlPng: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "pix", "png"),
			ImageUrlSvg: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "pix", "svg"),
		},
	}

	return m, nil
}

func (h *Handler) PromptPay(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payment := promptpay.PromptPay{
		PromptPayID: "0105540087061",   // Tax-ID/ID Card/E-Wallet
		Amount:      float64(m.Amount), // Positive amount
	}

	result, err := payment.Gen()
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "promptpay_display_qr_code",
		PromptpayDisplayQrCode: &next.PaymentIntentNextActionPromptpayDisplayQrCode{
			Data:        result,
			ImageUrlPng: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "promptpay", "png"),
			ImageUrlSvg: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(result)), "promptpay", "svg"),
		},
	}

	return m, nil
}
