package charge

import (
	"github.com/driver005/gateway/utils"
	"github.com/driver005/gateway/utils/review"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Charge

func (h *Handler) Bind(context *fiber.Ctx) (*Charge, error) {
	var m Charge
	var err error

	type request struct {
		*Alias
		Customer                  uuid.NullUUID `json:"customer,omitempty"`
		BalanceTransaction        uuid.NullUUID `json:"balance_transaction,omitempty"`
		FailureBalanceTransaction uuid.NullUUID `json:"failure_balance_transaction,omitempty"`
		PaymentMethod             uuid.NullUUID `json:"payment_method"`
		Review                    uuid.NullUUID `json:"review,omitempty"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context.Context(), model.Customer.UUID)
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
	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Review.Valid {
		m.Review, err = utils.Retrive[review.Review](h.r.Utils(), context.Context(), model.Review.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
