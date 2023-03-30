package credit

import (
	"github.com/gofiber/fiber/v2"
)

func (h *Handler) Bind(context *fiber.Ctx) (*CreditNote, error) {
	type Alias CreditNote
	var model CreditNote
	var err error

	if err = context.BodyParser(&model); err != nil {
		return nil, err
	}

	if model.CustomerReq.Valid {
		model.Customer, err = h.r.Customer().Retrive(context.Context(), model.CustomerReq.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.CustomerBalanceTransactionReq.Valid {
		model.CustomerBalanceTransaction, err = h.r.Balance().RetriveBalanceTransaction(context.Context(), model.CustomerBalanceTransactionReq.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.InvoiceReq.Valid {
		model.Invoice, err = h.r.Invoice().Retrive(context.Context(), model.InvoiceReq.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.RefundReq.Valid {
		model.Refund, err = h.r.Refund().Retrive(context.Context(), model.RefundReq.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &model, nil
}
