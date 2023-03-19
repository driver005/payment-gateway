package item

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*LineItem, error) {
	type Alias LineItem
	var m LineItem
	var err error

	model := struct {
		*Alias
		Product uuid.NullUUID `json:"product,omitempty"`
		Price   uuid.NullUUID `json:"price,omitempty"`
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
	if model.Price.Valid {
		m.Price, err = h.r.Price().Retrive(context.Context(), model.Price.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
