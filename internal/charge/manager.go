package charge

import (
	"github.com/driver005/gateway/utils"
	"github.com/driver005/gateway/utils/review"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Charge, error) {
	type Alias Charge
	var m Charge
	var err error

	model := struct {
		*Alias
		Customer                  uuid.NullUUID `json:"customer,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		BalanceTransaction        uuid.NullUUID `json:"balance_transaction,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		FailureBalanceTransaction uuid.NullUUID `json:"failure_balance_transaction,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		PaymentMethod             uuid.NullUUID `json:"payment_method" database:"default:null" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		Review                    uuid.NullUUID `json:"review,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
	}{
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
