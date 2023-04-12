package mandate

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Mandate, error) {
	type Alias Mandate
	var m Mandate
	var err error

	model := struct {
		*Alias
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty gorm:"-:all"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
