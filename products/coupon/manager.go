package coupon

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Coupon, error) {
	type Alias Coupon
	var m Coupon
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
