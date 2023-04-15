package setupIntent

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias SetupIntent

func (h *Handler) Bind(context *fiber.Ctx) (*SetupIntent, error) {
	var m SetupIntent
	var err error

	type request struct {
		*Alias
		Customer         uuid.NullUUID `json:"customer,omitempty"`
		PaymentMethod    uuid.NullUUID `json:"payment_method,omitempty"`
		Mandate          uuid.NullUUID `json:"mandate,omitempty"`
		SingleUseMandate uuid.NullUUID `json:"single_use_mandate,omitempty"`
	}

	var model = request{
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
