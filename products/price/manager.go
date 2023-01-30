package price

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Price, error) {
	type Alias Price
	var m Price
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
