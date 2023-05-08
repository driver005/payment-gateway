package invoiceItem

import (
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
		Customer         uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Invoice          uuid.NullUUID `json:"invoice,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Subscription     uuid.NullUUID `json:"subscription,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Price            uuid.NullUUID `json:"price,omitempty" swaggertype:"primitive,string" format:"uuid"`
		SubscriptionItem uuid.NullUUID `json:"subscription_item,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
	if model.Invoice.Valid {
		m.Invoice, err = h.r.Invoice().Retrive(context, model.Invoice.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Price.Valid {
		m.Price, err = h.r.Price().Retrive(context, model.Price.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.Subscription.Valid {
		m.Subscription, err = h.r.Subscription().Retrive(context, model.Subscription.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.SubscriptionItem.Valid {
		m.SubscriptionItem, err = h.r.SubscriptionItem().Retrive(context, model.SubscriptionItem.UUID)
		if err != nil {
			return nil, err
		}
	}

	return &m, nil
}

func (h *Handler) Manage(ctx *fiber.Ctx, m *Invoiceitem) {
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
