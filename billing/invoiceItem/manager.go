package invoiceItem

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Invoiceitem, error) {
	type Alias Invoiceitem
	var m Invoiceitem
	var err error

	model := struct {
		*Alias
		Customer     uuid.NullUUID `json:"customer,omitempty"`
		Invoice      uuid.NullUUID `json:"invoice,omitempty"`
		Subscription uuid.NullUUID `json:"subscription,omitempty"`
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
	if model.Invoice.Valid {
		m.Invoice, err = h.r.Invoice().Retrive(context.Context(), model.Invoice.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Subscription.Valid {
		m.Subscription, err = h.r.Subscription().Retrive(context.Context(), model.Subscription.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
