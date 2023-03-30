package quote

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Quote, error) {
	type Alias Quote
	var m Quote
	var err error

	model := struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		Invoice              uuid.NullUUID `json:"invoice,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		Subscription         uuid.NullUUID `json:"subscription,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		SubscriptionSchedule uuid.NullUUID `json:"subscription_schedule,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
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
	if model.SubscriptionSchedule.Valid {
		m.SubscriptionSchedule, err = h.r.SubscriptionSchedule().Retrive(context.Context(), model.SubscriptionSchedule.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
