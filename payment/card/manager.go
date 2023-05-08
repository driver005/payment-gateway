package card

import (
	"github.com/driver005/gateway/utils"
	"github.com/driver005/gateway/utils/account"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Card

func (h *Handler) Bind(context *fiber.Ctx) (*Card, error) {
	var m Card
	var err error

	type request struct {
		*Alias
		Account  uuid.NullUUID `json:"account,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Customer uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
	}

	var model = request{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Account.Valid {
		m.Account, err = utils.Retrive[account.Account](h.r.Utils(), context, model.Account.UUID)
		if err != nil {
			return nil, err
		}
	}

	if model.Customer.Valid {
		m.Customer, err = h.r.Customer().Retrive(context, model.Customer.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}
