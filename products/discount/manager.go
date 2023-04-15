package discount

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Discount

func (h *Handler) Bind(context *fiber.Ctx) (*Discount, error) {
	var m Discount
	var err error

	type request struct {
		*Alias
		Customer      uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Coupon        uuid.NullUUID `json:"coupon,omitempty" swaggertype:"primitive,string" format:"uuid"`
		PromotionCode uuid.NullUUID `json:"promotion_code,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
