package blik

import (
	"github.com/driver005/gateway/internal/intent"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Blik struct {
	ApiKey string
}

func NewBlik(apiKey string) *Blik {
	return &Blik{
		ApiKey: apiKey,
	}
}

func (c *Blik) Payment(m *intent.PaymentIntent) (*stripe.PaymentIntent, error) {
	stripe.Key = c.ApiKey

	params := &stripe.PaymentIntentParams{
		Confirm:            stripe.Bool(true),
		Amount:             stripe.Int64(1299),
		Currency:           stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{stripe.String("giropay")},
		ReturnURL:          stripe.String("http://localhost/complete"),
	}
	params.AddExtra("payment_method_data[type]", "giropay")
	params.AddExtra("payment_method_data[billing_details][name]", "Jenny Rosen")

	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
