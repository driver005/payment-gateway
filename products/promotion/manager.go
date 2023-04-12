package promotion

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*PromotionCode, error) {
	type Alias PromotionCode
	var m PromotionCode
	var err error

	model := struct {
		*Alias
		Customer uuid.NullUUID `json:"customer,omitempty gorm:"-:all"`
		Coupon   uuid.NullUUID `json:"coupon,omitempty gorm:"-:all"`
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
	if model.Coupon.Valid {
		m.Coupon, err = h.r.Coupon().Retrive(context.Context(), model.Coupon.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
