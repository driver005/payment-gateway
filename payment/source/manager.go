package source

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Source

func (h *Handler) Bind(context *fiber.Ctx) (*Source, error) {
	var m Source
	var err error

	type request struct {
		*Alias
		Customer uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
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

	return &m, nil
}
