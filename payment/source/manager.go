package source

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Source, error) {
	type Alias Source
	var m Source
	var err error

	model := struct {
		*Alias
		Customer uuid.NullUUID `json:"customer,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context.Context(), model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
