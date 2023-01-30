package item

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Bind(context *fiber.Ctx) (*LineItem, error) {
	type Alias LineItem
	var m LineItem
	var err error

	model := struct {
		*Alias
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	return &m, nil
}
