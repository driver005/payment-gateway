package invoiceItem

import (
	"context"
	"time"

	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/products/item"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type Alias Invoiceitem

func (h *Handler) Bind(context *fiber.Ctx) (*Invoiceitem, error) {
	var m Invoiceitem
	var err error

	type request struct {
		*Alias
		Customer         uuid.NullUUID `json:"customer,omitempty"`
		Invoice          uuid.NullUUID `json:"invoice,omitempty"`
		Subscription     uuid.NullUUID `json:"subscription,omitempty"`
		Price            uuid.NullUUID `json:"price,omitempty"`
		SubscriptionItem uuid.NullUUID `json:"subscription_item,omitempty"`
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
	if model.Invoice.Valid {
		m.Invoice, err = h.r.Invoice().Retrive(context.Context(), model.Invoice.UUID)
		if err != nil {
			return nil, err
		}
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
	if model.SubscriptionItem.Valid {
		m.SubscriptionItem, err = h.r.SubscriptionItem().Retrive(context.Context(), model.SubscriptionItem.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}

func (h *Handler) Manage(ctx context.Context, m *Invoiceitem) {
	i := invoice.Invoice{
		AmountDue:            m.Price.UnitAmount,
		AmountPaid:           m.Price.UnitAmount,
		Subtotal:             m.Price.UnitAmount,
		SubtotalExcludingTax: m.Price.UnitAmount - m.Invoice.Tax,
		Total:                m.Price.UnitAmount,
		TotalExcludingTax:    m.Price.UnitAmount - m.Invoice.Tax,
		Lines: []item.LineItem{
			{
				Amount:                 m.Price.UnitAmount,
				AmountExcludingTax:     m.Price.UnitAmount - m.Invoice.Tax,
				Description:            m.Price.Product.Description,
				Discountable:           m.Discountable,
				PeriodStart:            time.Now().UTC().Round(time.Second),
				PeriodEnd:              time.Now().UTC().Round(time.Second),
				Price:                  m.Price,
				UnitAmountExcludingTax: float64(m.Price.UnitAmount - m.Invoice.Tax),
				Quantity:               m.Quantity,
				Type:                   "invoiceitem",
			},
		},
	}

	i.Id = *m.InvoiceId

	h.r.Invoice().Update(ctx, &i)
}
