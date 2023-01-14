package driver

import (
	"github.com/driver005/gateway/driver/card"
	"github.com/stripe/stripe-go/v74"
)

func (h *Handler) Card() (*stripe.PaymentIntent, error) {
	c := card.NewCard("sk_test_51MJoIJDuTp4RMkBZ7gkfWY0Un5XLA60A4cpkn86lARUhCuflEj2BIN0iD80J7mRwZCQodp0azQ4BC5Wd71nZxOEF00g9d402Tk")
	payload, err := c.GeneratePayment()
	if err != nil {
		return nil, err
	}

	payment, err := c.ConfirmPayment(payload)
	if err != nil {
		return nil, err
	}

	return payment, nil
}
