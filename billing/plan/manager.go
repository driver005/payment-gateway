package plan

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Plan

func (h *Handler) Bind(context *fiber.Ctx) (*Plan, error) {
	var m Plan
	var err error

	type request struct {
		*Alias
		Product uuid.NullUUID `json:"product,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Product.Valid {
		m.Product, err = h.r.Product().Retrive(context.Context(), model.Product.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
