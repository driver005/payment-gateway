package driver

import (
	"github.com/driver005/gateway/driver/realtime/paynow"
	"github.com/driver005/gateway/driver/realtime/pix"
	"github.com/driver005/gateway/driver/realtime/promptpay"
)

func (h *Handler) PayNow() ([]byte, string, error) {
	payee := paynow.NewMobileQRCode("90991234")

	qrcode, result, err := payee.QRCode(12.35, "ABCDEFG")
	if err != nil {
		return nil, "", err
	}

	return qrcode, result, nil
}

func (h *Handler) Pix() ([]byte, string, error) {
	options := pix.Options{
		Name:          "Jonnas Fonini",
		Key:           "jonnasfonini@gmail.com",
		City:          "Marau",
		Amount:        20.67,        // optional
		Description:   "Invoice #4", // optional
		TransactionID: "***",        // optional
	}

	result, err := pix.Pix(options)
	if err != nil {
		return nil, "", err
	}

	qrcode, err := pix.QRCode(pix.QRCodeOptions{
		Content: result,
	})
	if err != nil {
		return nil, "", err
	}

	return qrcode, result, nil
}

func (h *Handler) PromptPay() ([]byte, string, error) {
	payment := promptpay.PromptPay{
		PromptPayID: "0105540087061", // Tax-ID/ID Card/E-Wallet
		Amount:      100.55,          // Positive amount
	}

	result, err := payment.Gen()
	if err != nil {
		return nil, "", err
	}

	qrcode, err := payment.QRCode(promptpay.QRCodeOptions{
		Content: result,
	})
	if err != nil {
		return nil, "", err
	}

	return qrcode, result, nil
}
