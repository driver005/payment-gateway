package driver

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image/png"
	"time"

	"github.com/driver005/gateway/driver/voucher/boleto"
	"github.com/driver005/gateway/driver/voucher/oxxo"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
)

func InitOxxo() *oxxo.OpClient {
	// publicApiKey := "key_EE7a1ZdR2eXtggex1dunMUw"
	privateApiKey := "key_qwPGiZ5pn4yOVBFlWRBYDEu"

	client := new(oxxo.OpClient)
	client.Init(privateApiKey)
	return client
}

func (h *Handler) Oxxo(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	_, err := h.oxxo.CreateOrder(oxxo.RequestData{
		"line_items": []oxxo.RequestData{
			{
				"name":       m.Description,
				"unit_price": m.Amount,
				"quantity":   1,
			},
		},
		"currency": "MXN",
		"customer_info": oxxo.RequestData{
			"name":  m.PaymentMethod.BillingDetails.Name,
			"email": m.PaymentMethod.BillingDetails.Email,
			"phone": m.PaymentMethod.BillingDetails.Phone,
			"charges": []oxxo.RequestData{
				{
					"payment_method": oxxo.RequestData{
						"type": "oxxo_cash",
					},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (h *Handler) Boleto(m *intent.PaymentIntent) *intent.PaymentIntent {
	var bytesPng bytes.Buffer

	// static data, you should keep this configured in somewhere
	var bank boleto.Bank = boleto.BB{
		Account:  88888888,
		Agency:   4444,
		Carteira: 55,
		Convenio: 4321,
		Company: boleto.Company{
			Name:      "Nome da empresa",
			LegalName: "Razao social",
			Address:   "Endereço",
			Contact:   "Email e telefone",
			Document:  "CNPJ/CPF",
		},
	}

	// dynamic data, you should have this data coming from a database
	var document = boleto.Document{
		Id:            1111,
		Value:         99999,
		ValueTax:      100,
		ValueDiscount: 0,
		ValueForfeit:  19999,
		OurNumber:     111111,
		FebrabanType:  "md",
		Date:          time.Now(),
		DateDue:       time.Now().AddDate(0, 0, 4),
		Instructions: [6]string{
			"Não receber após o vencimento",
			"Após vencimento, receber apenas no meu banco",
		},
		Payer: boleto.Payer{
			Name:    m.PaymentMethod.BillingDetails.Name,
			Address: m.PaymentMethod.BillingDetails.Address.Address1,
			Contact: m.PaymentMethod.BillingDetails.Email,
		},
	}

	// Optional, to use in your backend,
	// then you can save the barcode digitable number, or save the image separately
	var barcode boleto.Barcode = bank.Barcode(document)

	png.Encode(&bytesPng, barcode.Image())

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "boleto_display_details",
		BoletoDisplayDetails: &next.PaymentIntentNextActionBoleto{
			Number:           barcode.Digitable(),
			HostedVoucherUrl: fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(barcode.ToString())), "boleto", "png"),
			Pdf:              fmt.Sprintf("http://localhost:80/api/v1/code?data=%+v&paymentMethod=%+v&type=%+v", base64.URLEncoding.EncodeToString([]byte(barcode.ToString())), "boleto", "pdf"),
		},
	}

	return m

}
