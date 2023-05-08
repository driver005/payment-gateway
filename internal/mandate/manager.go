package mandate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Mandate

func (h *Handler) Bind(context *fiber.Ctx) (*Mandate, error) {
	var m Mandate
	var err error

	type request struct {
		*Alias
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context, model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
