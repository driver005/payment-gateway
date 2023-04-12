package bank

import (
	"github.com/driver005/gateway/utils"
	"github.com/driver005/gateway/utils/account"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (h *Handler) Bind(context *fiber.Ctx) (*BankAccount, error) {
	type Alias BankAccount
	var m BankAccount
	var err error

	model := struct {
		*Alias
		Account  uuid.NullUUID `json:"account,omitempty gorm:"-:all"`
		Customer uuid.NullUUID `json:"customer,omitempty gorm:"-:all"`
	}{
		Alias: (*Alias)(&m),
	}

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.Account.Valid {
		m.Account, err = utils.Retrive[account.Account](h.r.Utils(), context.Context(), model.Account.UUID)
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

	return &m, nil
}
