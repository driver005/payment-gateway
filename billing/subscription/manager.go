package subscription

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Subscription

func (h *Handler) Bind(context *fiber.Ctx) (*Subscription, error) {
	var m Subscription
	var err error

	type request struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		DefaultPaymentMethod uuid.NullUUID `json:"default_payment_method,omitempty" swaggertype:"primitive,string" format:"uuid"`
		PendingSetupIntent   uuid.NullUUID `json:"pending_setup_intent,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Schedule             uuid.NullUUID `json:"schedule,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Price                uuid.NullUUID `json:"price,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
	if model.DefaultPaymentMethod.Valid {
		m.DefaultPaymentMethod, err = h.r.PaymentMethod().Retrive(context, model.DefaultPaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PendingSetupIntent.Valid {
		m.PendingSetupIntent, err = h.r.SetupIntent().Retrive(context, model.PendingSetupIntent.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Price.Valid {
		m.Price, err = h.r.Price().Retrive(context, model.Price.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Schedule.Valid {
		m.Schedule, err = h.r.SubscriptionSchedule().Retrive(context, model.Schedule.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
