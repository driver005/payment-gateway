package attempt

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias SetupAttempt

func (h *Handler) Bind(context *fiber.Ctx) (*SetupAttempt, error) {
	var m SetupAttempt
	var err error

	type request struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty" swaggertype:"primitive,string" format:"uuid"`
		SetupIntent   uuid.NullUUID `json:"setup_intent,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context, model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.SetupIntent.Valid {
		m.SetupIntent, err = h.r.SetupIntent().Retrive(context, model.SetupIntent.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
