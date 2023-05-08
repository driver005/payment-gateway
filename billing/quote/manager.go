package quote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Quote

func (h *Handler) Bind(context *fiber.Ctx) (*Quote, error) {
	var m Quote
	var err error

	type request struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Invoice              uuid.NullUUID `json:"invoice,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Subscription         uuid.NullUUID `json:"subscription,omitempty" swaggertype:"primitive,string" format:"uuid"`
		SubscriptionSchedule uuid.NullUUID `json:"subscription_schedule,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
	if model.Invoice.Valid {
		m.Invoice, err = h.r.Invoice().Retrive(context, model.Invoice.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Subscription.Valid {
		m.Subscription, err = h.r.Subscription().Retrive(context, model.Subscription.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.SubscriptionSchedule.Valid {
		m.SubscriptionSchedule, err = h.r.SubscriptionSchedule().Retrive(context, model.SubscriptionSchedule.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
