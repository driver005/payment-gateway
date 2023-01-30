package attempt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*SetupAttempt, error) {
	type Alias SetupAttempt
	var m SetupAttempt
	var err error

	model := struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty"`
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty"`
		SetupIntent   uuid.NullUUID `json:"setup_intent,omitempty"`
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
	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.SetupIntent.Valid {
		m.SetupIntent, err = h.r.SetupIntent().Retrive(context.Context(), model.SetupIntent.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
