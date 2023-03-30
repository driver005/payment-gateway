package intent

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*SetupIntent, error) {
	type Alias SetupIntent
	var m SetupIntent
	var err error

	model := struct {
		*Alias
		Customer         uuid.NullUUID `json:"customer,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		PaymentMethod    uuid.NullUUID `json:"payment_method,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		Mandate          uuid.NullUUID `json:"mandate,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
		SingleUseMandate uuid.NullUUID `json:"single_use_mandate,omitempty" gorm:"-:all" swaggertype:"primitive,string" format:"uuid"`
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
	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Mandate.Valid {
		m.Mandate, err = h.r.Mandate().Retrive(context.Context(), model.Mandate.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.SingleUseMandate.Valid {
		m.SingleUseMandate, err = h.r.Mandate().Retrive(context.Context(), model.SingleUseMandate.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
