package product

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Product, error) {
	type Alias Product
	var m Product
	var err error

	model := struct {
		*Alias
		DefaultPrice uuid.NullUUID `json:"default_price,omitempty"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.DefaultPrice.Valid {
		m.DefaultPrice, err = h.r.Price().Retrive(context.Context(), model.DefaultPrice.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
