package intent

import (
	"context"

	"github.com/driver005/gateway/internal/intent/next"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx, pay bool, isPayable bool) (*PaymentIntent, error) {
	type Alias PaymentIntent
	var m PaymentIntent
	var err error

	model := struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty"`
		PaymentMethod uuid.NullUUID `json:"payment_method,omitempty"`
		Confirm       bool          `json:"confirm,omitempty"`
		ReturnUrl     string        `json:"return_url,omitempty"`
		OffSession    bool          `json:"off_session,omitempty"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.PaymentMethod.Valid {
		m.PaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.PaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context.Context(), model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}

	if model.ReturnUrl != "" {
		m.NextAction = &next.PaymentIntentNextAction{
			RedirectToUrl: &next.PaymentIntentNextActionRedirectToUrl{
				ReturnUrl: model.ReturnUrl,
			},
		}
	}
	if model.OffSession {
		m.SetupFutureUsage = "off_session"
	}

	if isPayable {
		if pay || model.Confirm {
			r, err := h.Pay(context.Context(), m)
			if err != nil {
				return nil, err
			}

			m = *r
		}
	}

	return &m, nil
}

func (h *Handler) Pay(ctx context.Context, m PaymentIntent) (*PaymentIntent, error) {
	r, err := h.r.Pay(ctx, &m)
	if err != nil {
		return nil, err
	}

	return r, nil
}
