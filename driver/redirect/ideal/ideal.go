package ideal

import (
	"github.com/driver005/gateway/internal/intent"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Ideal struct {
	ApiKey string
}

func NewIdeal(apiKey string) *Ideal {
	return &Ideal{
		ApiKey: apiKey,
	}
}

func (c *Ideal) Payment(m *intent.PaymentIntent) (*stripe.PaymentIntent, error) {
	stripe.Key = c.ApiKey

	params := &stripe.PaymentIntentParams{
		Confirm:            stripe.Bool(true),
		Amount:             stripe.Int64(1299),
		Currency:           stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{stripe.String("ideal")},
		ReturnURL:          stripe.String("http://localhost/complete"),
	}
	params.AddExtra("payment_method_data[type]", "ideal")

	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
