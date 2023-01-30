package subscription

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Subscription, error) {
	type Alias Subscription
	var m Subscription
	var err error

	model := struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty"`
		DefaultPaymentMethod uuid.NullUUID `json:"default_payment_method,omitempty"`
		PendingSetupIntent   uuid.NullUUID `json:"pending_setup_intent,omitempty"`
		Schedule             uuid.NullUUID `json:"schedule,omitempty"`
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
	if model.DefaultPaymentMethod.Valid {
		m.DefaultPaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.DefaultPaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PendingSetupIntent.Valid {
		m.PendingSetupIntent, err = h.r.SetupIntent().Retrive(context.Context(), model.PendingSetupIntent.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Schedule.Valid {
		m.Schedule, err = h.r.SubscriptionSchedule().Retrive(context.Context(), model.Schedule.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
