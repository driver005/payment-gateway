package price

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Price, error) {
	type Alias Price
	var m Price
	var err error

	model := struct {
		*Alias
		Product uuid.NullUUID `json:"product,omitempty"`
	}{
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
