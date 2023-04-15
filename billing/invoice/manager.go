package invoice

import (
	"context"
	"time"

	"github.com/driver005/gateway/internal/intent"
	"github.com/driver005/gateway/products/item"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Alias Invoice

func (h *Handler) Bind(context *fiber.Ctx) (*Invoice, error) {
	var m Invoice
	var err error

	type request struct {
		*Alias
		Customer             uuid.NullUUID `json:"customer,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Charge               uuid.NullUUID `json:"charge,omitempty" swaggertype:"primitive,string" format:"uuid"`
		DefaultPaymentMethod uuid.NullUUID `json:"default_payment_method,omitempty" swaggertype:"primitive,string" format:"uuid"`
		PaymentIntent        uuid.NullUUID `json:"payment_intent,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Price                uuid.NullUUID `json:"price,omitempty" swaggertype:"primitive,string" format:"uuid"`
		LatestRevision       uuid.NullUUID `json:"latest_revision,omitempty" swaggertype:"primitive,string" format:"uuid"`
		Subscription         uuid.NullUUID `json:"subscription,omitempty" swaggertype:"primitive,string" format:"uuid"`
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
	if model.Charge.Valid {
		m.Charge, err = h.r.Charge().Retrive(context.Context(), model.Charge.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.DefaultPaymentMethod.Valid {
		m.DefaultPaymentMethod, err = h.r.PaymentMethod().Retrive(context.Context(), model.DefaultPaymentMethod.UUID)
		if err != nil {
			return nil, err
		}
	}
	if model.PaymentIntent.Valid {
		m.PaymentIntent, err = h.r.PaymentIntent().Retrive(context.Context(), model.PaymentIntent.UUID)
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
	if model.LatestRevision.Valid {
		m.LatestRevision, err = h.Retrive(context.Context(), model.LatestRevision.UUID)
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

	return h.Manage(context.Context(), m)
}

func (h *Handler) Manage(ctx context.Context, m Invoice) (*Invoice, error) {
	var paymentIntent *intent.PaymentIntent
	var err error

	billingReason := m.BillingReason
	collectionMethod := m.CollectionMethod
	dueDate := m.DueDate
	status := m.Status
	now := time.Now().UTC().Round(time.Second)

	if billingReason == "" {
		billingReason = BillingReasonManual
	}
	if collectionMethod == "" {
		collectionMethod = CollectionMethodSendInvoiced
	}
	if dueDate.IsZero() {
		dueDate = time.Now().AddDate(0, 0, 30)
	}
	if status == "" {
		status = StatusDraft
	}

	if status == "open" {
		paymentIntent, err = h.r.PaymentIntent().Create(
			ctx,
			&intent.PaymentIntent{
				Customer:    m.Customer,
				Amount:      m.Price.UnitAmount,
				Description: "PaymentIntent for Invoice",
				PaymentMethodTypes: pq.StringArray{
					"bancontact",
					"card",
					"eps",
					"giropay",
					"ideal",
					"p24",
					"sepa_debit",
					"sofort",
					"wechat_pay",
				},
			},
		)
		if err != nil {
			return nil, err
		}
	}

	m.AmountDue = m.Price.UnitAmount
	m.AmountPaid = 0
	m.AmountRemaining = m.Price.UnitAmount
	m.DueDate = dueDate
	m.BillingReason = billingReason
	m.CollectionMethod = collectionMethod
	m.CustomerAddress = m.Customer.Address
	m.CustomerEmail = m.Customer.Email
	m.CustomerName = m.Customer.Name
	m.CustomerPhone = m.Customer.Phone
	m.CustomerShipping = m.Customer.Shipping
	m.CustomerTaxExempt = m.Customer.TaxExempt
	m.Lines = []item.LineItem{
		{
			Amount:                 m.Price.UnitAmount,
			AmountExcludingTax:     m.Price.UnitAmount - m.Tax,
			Description:            m.Price.Product.Description,
			Discountable:           m.Discountable,
			PeriodStart:            now,
			PeriodEnd:              now,
			Price:                  m.Price,
			UnitAmountExcludingTax: float64(m.Price.UnitAmount - m.Tax),
			Quantity:               m.Quantity,
			Type:                   "invoice",
		},
	}
	m.PeriodStart = now
	m.PeriodEnd = now
	m.Subtotal = m.Price.UnitAmount
	m.SubtotalExcludingTax = m.Price.UnitAmount - m.Tax
	m.Total = m.Price.UnitAmount
	m.TotalExcludingTax = m.Price.UnitAmount - m.Tax
	m.Status = status
	m.PaymentIntent = paymentIntent

	return &m, nil
}
