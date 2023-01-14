package card

import (
	"github.com/stripe/stripe-go/v74"
	"github.com/stripe/stripe-go/v74/paymentintent"
	"github.com/stripe/stripe-go/v74/paymentmethod"
)

type Card struct {
	ApiKey string
}

func NewCard(apiKey string) Card {
	return Card{
		ApiKey: apiKey,
	}
}

func (c *Card) GeneratePayment() (*stripe.PaymentIntent, error) {
	stripe.Key = c.ApiKey

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(1299),
		Currency:           stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: []*string{stripe.String("card")},
	}
	result, err := paymentintent.New(params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Card) GeneratePaymentMethod() (*stripe.PaymentMethod, error) {
	params := stripe.PaymentMethodParams{
		Card: &stripe.PaymentMethodCardParams{
			Number:   stripe.String("4242424242424242"),
			ExpMonth: stripe.Int64(8),
			ExpYear:  stripe.Int64(2024),
			CVC:      stripe.String("314"),
		},
		Type: stripe.String("card"),
	}

	result, err := paymentmethod.New(&params)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (c *Card) ConfirmPayment(p *stripe.PaymentIntent) (*stripe.PaymentIntent, error) {
	paymentMethod, err := c.GeneratePaymentMethod()
	if err != nil {
		return nil, err
	}

	params := &stripe.PaymentIntentConfirmParams{
		PaymentMethod: &paymentMethod.ID,
	}

	result, err := paymentintent.Confirm(p.ID, params)
	if err != nil {
		return nil, err
	}

	return result, nil
}
