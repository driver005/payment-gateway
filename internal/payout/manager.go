package payout

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Payout

func (h *Handler) Bind(context *fiber.Ctx) (*Payout, error) {
	var m Payout
	var err error

	type request struct {
		*Alias
		BalanceTransaction        uuid.NullUUID `json:"balance_transaction,omitempty" swaggertype:"primitive,string" format:"uuid"`
		FailureBalanceTransaction uuid.NullUUID `json:"failure_balance_transaction,omitempty" swaggertype:"primitive,string" format:"uuid"`
		OriginalPayout            uuid.NullUUID `json:"original_payout,omitempty" swaggertype:"primitive,string" format:"uuid"`
		ReversedBy                uuid.NullUUID `json:"reversed_by,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
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
	if model.OriginalPayout.Valid {
		m.OriginalPayout, err = h.Retrive(context.Context(), model.OriginalPayout.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.ReversedBy.Valid {
		m.ReversedBy, err = h.Retrive(context.Context(), model.ReversedBy.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
