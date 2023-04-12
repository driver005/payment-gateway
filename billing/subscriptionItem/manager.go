package subscriptionItem

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*SubscriptionItem, error) {
	type Alias SubscriptionItem
	var m SubscriptionItem
	var err error

	model := struct {
		*Alias
		Subscription uuid.NullUUID `json:"subscription,omitempty gorm:"-:all"`
		Price        uuid.NullUUID `json:"price,omitempty gorm:"-:all"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Price.Valid {
		m.Price, err = h.r.Price().Retrive(context.Context(), model.Price.UUID)
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
