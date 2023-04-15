package coupon

import (
	"github.com/gofiber/fiber/v2"
)

type Alias Coupon

func (h *Handler) Bind(context *fiber.Ctx) (*Coupon, error) {
	var m Coupon
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
