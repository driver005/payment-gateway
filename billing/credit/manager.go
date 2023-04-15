package credit

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias CreditNote

func (h *Handler) Bind(context *fiber.Ctx) (*CreditNote, error) {
	var m CreditNote
	var err error

	type request struct {
		*Alias
		Customer                   uuid.NullUUID `json:"customer,omitempty"`
		CustomerBalanceTransaction uuid.NullUUID `json:"customer_balance_transaction,omitempty"`
		Invoice                    uuid.NullUUID `json:"invoice,omitempty"`
		Refund                     uuid.NullUUID `json:"refund,omitempty"`
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
	if model.CustomerBalanceTransaction.Valid {
		m.CustomerBalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context.Context(), model.CustomerBalanceTransaction.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Invoice.Valid {
		m.Invoice, err = h.r.Invoice().Retrive(context.Context(), model.Invoice.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Refund.Valid {
		m.Refund, err = h.r.Refund().Retrive(context.Context(), model.Refund.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
