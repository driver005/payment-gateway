package promotion

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias PromotionCode

func (h *Handler) Bind(context *fiber.Ctx) (*PromotionCode, error) {
	var m PromotionCode
	var err error

	type request struct {
		*Alias
		Customer uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Coupon   uuid.NullUUID `json:"coupon,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context, model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Coupon.Valid {
		m.Coupon, err = h.r.Coupon().Retrive(context, model.Coupon.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
