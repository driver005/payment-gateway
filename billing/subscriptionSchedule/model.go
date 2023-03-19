package subscriptionSchedule

import (
	"time"

	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/products/coupon"
	"github.com/driver005/gateway/products/price"
	"github.com/driver005/gateway/utils/tax"
	"github.com/google/uuid"
)

// SubscriptionSchedulesResourceDefaultSettings
type SubscriptionSchedulesResourceDefaultSettings struct {
	core.Model

	AutomaticTax bool `json:"automatic_tax,omitempty"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor string `json:"billing_cycle_anchor,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod     string                `json:"collection_method,omitempty"`
	DefaultPaymentMethod *method.PaymentMethod `json:"default_payment_method,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description,omitempty"`
	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int `json:"days_until_due,omitempty"`
}

// SubscriptionScheduleAddInvoiceItem An Add Invoice Item describes the prices and quantities that will be added as pending invoice items when entering a phase.
type SubscriptionScheduleSubscriptionItem struct {
	core.Model

	Price *price.Price `json:"price" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The quantity of the invoice item.
	Quantity int `json:"quantity,omitempty"`
	// The tax rates which apply to the item. When set, the `default_tax_rates` do not apply to this item.
	TaxRates *tax.TaxRate `json:"tax_rates,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

// SubscriptionSchedulePhaseConfiguration A phase describes the plans, coupon, and trialing status of a subscription for a predefined time period.
type SubscriptionSchedulePhaseConfiguration struct {
	core.Model

	// A list of prices and quantities that will generate invoice items appended to the next invoice for this phase.
	AddInvoiceItems []SubscriptionScheduleSubscriptionItem `json:"add_invoice_items,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	AutomaticTax    bool                                   `json:"automatic_tax,omitempty"`
	// Possible values are `phase_start` or `automatic`. If `phase_start` then billing cycle anchor of the subscription is set to the start of the phase when entering the phase. If `automatic` then the billing cycle anchor is automatically modified as needed when entering the phase. For more information, see the billing cycle [documentation](https://stripe.com/docs/billing/subscriptions/billing-cycle).
	BillingCycleAnchor string `json:"billing_cycle_anchor,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay the underlying subscription at the end of each billing cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod string         `json:"collection_method,omitempty"`
	Coupon           *coupon.Coupon `json:"coupon,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency             string                `json:"currency,omitempty"`
	DefaultPaymentMethod *method.PaymentMethod `json:"default_payment_method,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The default tax rates to apply to the subscription during this phase of the subscription schedule.
	DefaultTaxRates *tax.TaxRate `json:"default_tax_rates,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Subscription description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription.
	Description string `json:"description,omitempty"`
	// The end of this phase of the subscription schedule.
	EndDate         int                                                 `json:"end_date,omitempty"`
	InvoiceSettings *settings.InvoiceSettingSubscriptionScheduleSetting `json:"invoice_settings,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Subscription items to configure the subscription to during this phase of the subscription schedule.
	Items []SubscriptionScheduleSubscriptionItem `json:"items,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// If the subscription schedule will prorate when transitioning to this phase. Possible values are `create_prorations` and `none`.
	ProrationBehavior string `json:"proration_behavior,omitempty"`
	// The start of this phase of the subscription schedule.
	StartDate int `json:"start_date,omitempty"`
	// When the trial ends within the phase.
	TrialEnd int `json:"trial_end,omitempty"`
}

// SubscriptionSchedule A subscription schedule allows you to create and manage the lifecycle of a subscription by predefining expected changes.  Related guide: [Subscription Schedules](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
type SubscriptionSchedule struct {
	core.Model

	// Time at which the subscription schedule was canceled. Measured in seconds since the Unix epoch.
	CanceledAt int `json:"canceled_at,omitempty"`
	// Time at which the subscription schedule was completed. Measured in seconds since the Unix epoch.
	CompletedAt int `json:"completed_at,omitempty"`
	// The end of this phase of the subscription schedule.
	EndDate int `json:"end_date,omitempty"`
	// The start of this phase of the subscription schedule.
	StartDate       int                                           `json:"start_date,omitempty"`
	DefaultSettings *SubscriptionSchedulesResourceDefaultSettings `json:"default_settings,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Behavior of the subscription schedule and underlying subscription when it ends. Possible values are `release` or `cancel` with the default being `release`. `release` will end the subscription schedule and keep the underlying subscription running.`cancel` will end the subscription schedule and cancel the underlying subscription.
	EndBehavior string `json:"end_behavior,omitempty"`
	// Configuration for the subscription schedule's phases.
	Phases *SubscriptionSchedulePhaseConfiguration `json:"phases,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Time at which the subscription schedule was released. Measured in seconds since the Unix epoch.
	ReleasedAt time.Time `json:"released_at,omitempty"`
	// ID of the subscription once managed by the subscription schedule (if it is released).
	ReleasedSubscription *uuid.UUID `json:"released_subscription,omitempty"`
	// The present status of the subscription schedule. Possible values are `not_started`, `active`, `completed`, `released`, and `canceled`. You can read more about the different states in our [behavior guide](https://stripe.com/docs/billing/subscriptions/subscription-schedules).
	Status string `json:"status,omitempty"`

	CustomerId *uuid.UUID         `json:"customer_id,omitempty" swaggerignore:"true"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
}
