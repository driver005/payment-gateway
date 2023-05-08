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
		Customer                  uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		BalanceTransaction        uuid.NullUUID `json:"balance_transaction,omitempty" swaggertype:"primitive,string" format:"uuid"`
		FailureBalanceTransaction uuid.NullUUID `json:"failure_balance_transaction,omitempty" swaggertype:"primitive,string" format:"uuid"`
		PaymentMethod             uuid.NullUUID `json:"payment_method" swaggertype:"primitive,string" format:"uuid"`
		Review                    uuid.NullUUID `json:"review,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context, model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.BalanceTransaction.Valid {
		m.BalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context, model.BalanceTransaction.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.FailureBalanceTransaction.Valid {
		m.FailureBalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context, model.FailureBalanceTransaction.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context, model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Review.Valid {
		m.Review, err = utils.Retrive[review.Review](h.r.Utils(), context, model.Review.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
