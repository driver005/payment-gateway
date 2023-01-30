package driver

import (
	"github.com/driver005/gateway/driver/card"
	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/spf13/viper"
)

func InitCard() *card.Card {
	return card.NewCard(viper.GetString("stripe.secret_key"))
}

func (h *Handler) Card(m *intent.PaymentIntent) (*intent.PaymentIntent, error) {
	payload, err := h.card.GeneratePayment(m)
	if err != nil {
		return nil, err
	}

	payment, err := h.card.ConfirmPayment(m, payload)
	if err != nil {
		return nil, err
	}

	m.NextAction = &next.PaymentIntentNextAction{
		Type: "card_await_notification",
		CardAwaitNotification: &next.PaymentIntentNextActionCardAwaitNotification{
			ChargeAttemptAt:          int(payment.NextAction.CardAwaitNotification.ChargeAttemptAt),
			CustomerApprovalRequired: payment.NextAction.CardAwaitNotification.CustomerApprovalRequired,
		},
	}

	return m, nil
}
