package customer

import (
	"github.com/driver005/gateway/helper"
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Customer, error) {
	var m Customer
	var err error

	if err = context.BodyParser(&m); err != nil {
		return nil, err
	}

	if m.InvoicePrefix == "" {
		m.InvoicePrefix = helper.String(8)
	}

	return &m, nil
}
