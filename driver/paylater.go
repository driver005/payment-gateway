package driver

import (
	"net/url"
	"strings"
	"time"

	"github.com/driver005/gateway/driver/paylater/affirm"
	"github.com/driver005/gateway/driver/paylater/klarna"
	"github.com/driver005/gateway/driver/paylater/riverty"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
)

//
// Affirm
//

func InitAffirm() *affirm.Affirm {
	uri, _ := url.Parse(affirm.API)
	conf := affirm.Config{
		BaseURL: uri,
		APIKey:  "",
		Timeout: time.Second * 10,
	}

	client := affirm.NewClient(conf)

	cl := affirm.NewAffirm(client)

	return &cl
}

func (h *Handler) Affirm(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	resp, err := h.affirm.CreateDirectCheckout(&affirm.DirectCheckout{
		Billing: affirm.Billing{
			Name: affirm.Name{
				First: strings.Split(m.PaymentMethod.BillingDetails.Name, " ")[0],
				Last:  strings.Split(m.PaymentMethod.BillingDetails.Name, " ")[1],
				Full:  m.PaymentMethod.BillingDetails.Name,
			},
			Address: affirm.Address{
				City:    m.PaymentMethod.BillingDetails.Address.City,
				Country: m.PaymentMethod.BillingDetails.Address.Country.DisplayName,
				Line1:   m.PaymentMethod.BillingDetails.Address.Address1,
				Line2:   m.PaymentMethod.BillingDetails.Address.Address2,
				State:   m.PaymentMethod.BillingDetails.Address.Province,
				Zipcode: m.PaymentMethod.BillingDetails.Address.PostalCode,
			},
			PhoneNumber: m.PaymentMethod.BillingDetails.Phone,
			Email:       m.PaymentMethod.BillingDetails.Email,
		},
		Total:     m.Amount,
		OrderId:   m.Id.String(),
		TaxAmount: 0,
		Shipping: affirm.Shipping{
			Name: affirm.Name{
				First: strings.Split(m.Shipping.Name, " ")[0],
				Last:  strings.Split(m.Shipping.Name, " ")[1],
				Full:  m.Shipping.Name,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{

			Url: resp.RedirectUrl,
		},
	}

	return m, nil
}

//
// Riverty
//

func InitRiverty() *riverty.Riverty {
	uri, _ := url.Parse(riverty.API)
	conf := riverty.Config{
		BaseURL: uri,
		APIKey:  "",
		Timeout: time.Second * 10,
	}

	client := riverty.NewClient(conf)

	cl := riverty.NewRiverty(client)

	return &cl
}

func (h *Handler) Riverty(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	resp, err := h.riverty.CreateCheckout(&riverty.Checkout{
		CheckoutID: m.Id.String(),
		Payment: riverty.Payment{
			Type: m.PaymentMethod.AfterpayClearpay.Type,
		},
		Customer: riverty.Customer{
			Address: riverty.Address{
				Street:       m.PaymentMethod.BillingDetails.Address.Address1,
				StreetNumber: m.PaymentMethod.BillingDetails.Address.Address1,
				PostalCode:   m.PaymentMethod.BillingDetails.Address.PostalCode,
				PostalPlace:  m.PaymentMethod.BillingDetails.Address.City,
				CountryCode:  m.PaymentMethod.BillingDetails.Address.CountryCode,
			},
			RiskData: riverty.RiskData{
				IPAddress: "",
			},
			FirstName:        strings.Split(m.PaymentMethod.BillingDetails.Name, " ")[0],
			LastName:         strings.Split(m.PaymentMethod.BillingDetails.Name, " ")[1],
			Email:            m.PaymentMethod.BillingDetails.Email,
			BirthDate:        "",
			CustomerCategory: "Person",
		},
		Order: riverty.Order{
			TotalNetAmount:   m.Amount,
			TotalGrossAmount: m.Amount,
			Currency:         m.Currency,
			Items: []riverty.Items{
				{
					ProductID:      "",
					Description:    m.Description,
					NetUnitPrice:   m.Amount,
					GrossUnitPrice: m.Amount,
					Quantity:       1,
					VatPercent:     0,
					VatAmount:      0,
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{

			Url: resp.SecureLoginURL,
		},
	}

	return m, nil
}

//
// Klarna
//

func InitKlarna() *klarna.PaymentSrv {
	uri, _ := url.Parse(riverty.API)
	conf := klarna.Config{
		BaseURL:     uri,
		APIUsername: "",
		APIPassword: "",
		Timeout:     time.Second * 10,
	}

	client := klarna.NewClient(conf)

	cl := klarna.NewPaymentSrv(client)

	return &cl
}

func (h *Handler) Klarna(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	resp, err := h.klarna.CreateNewOrder(m.Id.String(), &klarna.PaymentOrder{
		PurchaseCountry:  m.PaymentMethod.BillingDetails.Address.Country.Iso2,
		PurchaseCurrency: m.Currency,
		Locale:           "en-US",
		OrderAmount:      m.Amount,
		OrderTaxAmount:   0,
		OrderLines: []*klarna.Line{
			{
				Name:           m.Description,
				Quantity:       1,
				UnitPrice:      m.Amount,
				TaxRate:        0,
				TotalAmount:    m.Amount,
				TotalTaxAmount: 0,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "redirect_to_url",
		RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{

			Url: resp.RedirectURL,
		},
	}

	return m, nil
}
