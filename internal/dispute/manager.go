package dispute

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Dispute

func (h *Handler) Bind(context *fiber.Ctx) (*Dispute, error) {
	var m Dispute
	var err error

	type request struct {
		*Alias
		Charge        uuid.NullUUID `json:"charge,omitempty"`
		PaymentIntent uuid.NullUUID `json:"payment_intent,omitempty"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Charge.Valid {
		m.Charge, err = h.r.Charge().Retrive(context.Context(), model.Charge.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PaymentIntent.Valid {
		m.PaymentIntent, err = h.r.PaymentIntent().Retrive(context.Context(), model.PaymentIntent.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
