package invoice

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Invoice, error) {
	type Alias Invoice
	var m Invoice
	var err error

	model := struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty"`
		Charge               uuid.NullUUID `json:"charge,omitempty"`
		DefaultPaymentMethod uuid.NullUUID `json:"default_payment_method,omitempty"`
		PaymentIntent        uuid.NullUUID `json:"payment_intent,omitempty"`
		LatestRevision       uuid.NullUUID `json:"latest_revision,omitempty"`
		Subscription         uuid.NullUUID `json:"subscription,omitempty"`
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
	if model.Charge.Valid {
		m.Charge, err = h.r.Charge().Retrive(context.Context(), model.Charge.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.DefaultPaymentMethod.Valid {
		m.DefaultPaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.DefaultPaymentMethod.UUID)
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
	if model.LatestRevision.Valid {
		m.LatestRevision, err = h.Retrive(context.Context(), model.LatestRevision.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Subscription.Valid {
		m.Subscription, err = h.r.Subscription().Retrive(context.Context(), model.Subscription.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
