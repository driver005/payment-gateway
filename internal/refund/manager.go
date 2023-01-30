package refund

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Refund, error) {
	type Alias Refund
	var m Refund
	var err error

	model := struct {
		*Alias
		Charge                    uuid.NullUUID `json:"charge,omitempty"`
		PaymentIntent             uuid.NullUUID `json:"payment_intent,omitempty"`
		BalanceTransaction        uuid.NullUUID `json:"balance_transaction,omitempty"`
		FailureBalanceTransaction uuid.NullUUID `json:"failure_balance_transaction,omitempty"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Charge.Valid {
		m.Charge, err = h.r.Charge().Retrive(context.Context(), model.Charge.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PaymentIntent.Valid {
		m.PaymentIntent, err = h.r.PaymentIntent().Retrive(context.Context(), model.PaymentIntent.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.BalanceTransaction.Valid {
		m.BalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context.Context(), model.BalanceTransaction.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.FailureBalanceTransaction.Valid {
		m.FailureBalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context.Context(), model.FailureBalanceTransaction.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
