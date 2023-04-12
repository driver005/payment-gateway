package discount

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*Discount, error) {
	type Alias Discount
	var m Discount
	var err error

	model := struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty gorm:"-:all"`
		Coupon        uuid.NullUUID `json:"coupon,omitempty gorm:"-:all"`
		PromotionCode uuid.NullUUID `json:"promotion_code,omitempty gorm:"-:all"`
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
	if model.PromotionCode.Valid {
		m.PromotionCode, err = h.r.Promotion().Retrive(context.Context(), model.PromotionCode.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
