package product

import (
	"github.com/gofiber/fiber/v2"
)

type Alias Product

func (h *Handler) Bind(context *fiber.Ctx) (*Product, error) {
	var m Product
	var err error

	type request struct {
		*Alias
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	return &m, nil
}
