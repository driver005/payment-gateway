package eps

import (
	"github.com/driver005/gateway/internal/intent"
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
)

type Eps struct {
	ApiKey string
}

func NewEps(apiKey string) *Eps {
	return &Eps{
		ApiKey: apiKey,
	}
}

func (c *Eps) Payment(m *intent.PaymentIntent) (*stripe.PaymentIntent, error) {
	stripe.Key = c.ApiKey

	params := &stripe.PaymentIntentParams{
		Confirm:            stripe.Bool(true),
		Amount:             stripe.Int64(1299),
		Currency:           stripe.String(string(stripe.CurrencyEUR)),
		PaymentMethodTypes: []*string{stripe.String("eps")},
		ReturnURL:          stripe.String("http://localhost/complete"),
	}
	params.AddExtra("payment_method_data[type]", "eps")
	params.AddExtra("payment_method_data[eps][bank]", "bank_austria")

	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
