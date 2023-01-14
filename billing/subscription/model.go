package subscription

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/subscription/options"
	"github.com/driver005/gateway/billing/subscriptionItem"
	"github.com/driver005/gateway/billing/subscriptionSchedule"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/payment/method"
)

// SubscriptionsResourcePaymentSettings
type SubscriptionsResourcePaymentSettings struct {
	PaymentMethodOptions options.SubscriptionsResourcePaymentMethodOptions `json:"payment_method_options,omitempty"`
	// The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes []string `json:"payment_method_types,omitempty"`
	// Either `off`, or `on_subscription`. With `on_subscription` Stripe updates `subscription.default_payment_method` when a subscription payment succeeds.
	SaveDefaultPaymentMethod string `json:"save_default_payment_method,omitempty"`
}

// SubscriptionsResourcePauseCollection The Pause Collection settings determine how we will pause collection for this subscription and for how long the subscription should be paused.
type SubscriptionsResourcePauseCollection struct {
	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior string `json:"behavior"`
	// The time after which the subscription will resume collecting payments.
	ResumesAt int `json:"resumes_at,omitempty"`
}

// SubscriptionsResourcePendingUpdate Pending Updates store the changes pending from a previous update that will be applied to the Subscription upon successful payment.
type SubscriptionsResourcePendingUpdate struct {
	// If the update is applied, determines the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.
	BillingCycleAnchor int `json:"billing_cycle_anchor,omitempty"`
	// The point after which the changes reflected by this update will be discarded and no longer applied.
	ExpiresAt int32 `json:"expires_at"`
	// List of subscription items, each with an attached plan, that will be set if the update is applied.
	SubscriptionItems []subscriptionItem.SubscriptionItem `json:"subscription_items,omitempty"`
	// Unix timestamp representing the end of the trial period the customer will get before being charged for the first time, if the update is applied.
	TrialEnd int `json:"trial_end,omitempty"`
	// Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.
	TrialFromPlan bool `json:"trial_from_plan,omitempty"`
}

// SubscriptionBillingThresholds
type SubscriptionBillingThresholds struct {
	// Monetary threshold that triggers the subscription to create an invoice
	AmountGte int `json:"amount_gte,omitempty"`
	// Indicates if the `billing_cycle_anchor` should be reset when a threshold is reached. If true, `billing_cycle_anchor` will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be `true` if the subscription contains items with plans that have `aggregate_usage=last_ever`.
	ResetBillingCycleAnchor bool `json:"reset_billing_cycle_anchor,omitempty"`
}

// SubscriptionAutomaticTax
type SubscriptionAutomaticTax struct {
	// Whether Stripe automatically computes tax on this subscription.
	Enabled bool `json:"enabled"`
}

// SubscriptionPendingInvoiceItemInterval
type SubscriptionPendingInvoiceItemInterval struct {
	// Specifies invoicing frequency. Either `day`, `week`, `month` or `year`.
	Interval string `json:"interval"`
	// The number of intervals between invoices. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount int32 `json:"interval_count"`
}

// SubscriptionDefaultSource ID of the default payment source for the subscription. It must belong to the customer associated with the subscription and be in a chargeable state. If `default_payment_method` is also set, `default_payment_method` will take precedence. If neither are set, invoices will use the customer's [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/object#customer_object-invoice_settings-default_payment_method) or [default_source](https://stripe.com/docs/api/customers/object#customer_object-default_source).
type SubscriptionDefaultSource struct {
	// BankAccount *BankAccount
	// Card *Card
	// Source *Source
}

// Subscription Subscriptions allow you to charge a customer on a recurring basis.  Related guide: [Creating Subscriptions](https://stripe.com/docs/billing/subscriptions/creating).
type Subscription struct {
	AutomaticTax SubscriptionAutomaticTax `json:"automatic_tax"`
	// Determines the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.
	BillingCycleAnchor int32                         `json:"billing_cycle_anchor"`
	BillingThresholds  SubscriptionBillingThresholds `json:"billing_thresholds,omitempty"`
	// A date in the future at which the subscription will automatically get canceled
	CancelAt int `json:"cancel_at,omitempty"`
	// If the subscription has been canceled with the `at_period_end` flag set to `true`, `cancel_at_period_end` on the subscription will be true. You can use this attribute to determine whether a subscription that has a status of active is scheduled to be canceled at the end of the current period.
	CancelAtPeriodEnd bool `json:"cancel_at_period_end"`
	// If the subscription has been canceled, the date of that cancellation. If the subscription was canceled with `cancel_at_period_end`, `canceled_at` will reflect the time of the most recent update request, not the end of the subscription period when the subscription is automatically moved to a canceled state.
	CanceledAt int `json:"canceled_at,omitempty"`
	// Either `charge_automatically`, or `send_invoice`. When charging automatically, Stripe will attempt to pay this subscription at the end of the cycle using the default source attached to the customer. When sending an invoice, Stripe will email your customer an invoice with payment instructions and mark the subscription as `active`.
	CollectionMethod string `json:"collection_method"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int32 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// End of the current period that the subscription has been invoiced for. At the end of this period, a new invoice will be created.
	CurrentPeriodEnd int32 `json:"current_period_end"`
	// Start of the current period that the subscription has been invoiced for.
	CurrentPeriodStart int32             `json:"current_period_start"`
	Customer           customer.Customer `json:"customer"`
	// Number of days a customer has to pay invoices generated by this subscription. This value will be `null` for subscriptions where `collection_method=charge_automatically`.
	DaysUntilDue         int                       `json:"days_until_due,omitempty"`
	DefaultPaymentMethod method.PaymentMethod      `json:"default_payment_method,omitempty"`
	DefaultSource        SubscriptionDefaultSource `json:"default_source,omitempty"`
	// The tax rates that will apply to any subscription item that does not have `tax_rates` set. Invoices created will have their `default_tax_rates` populated from the subscription.
	DefaultTaxRates []models.TaxRate `json:"default_tax_rates,omitempty"`
	// The subscription's description, meant to be displayable to the customer. Use this field to optionally store an explanation of the subscription for rendering in Stripe surfaces.
	Description string `json:"description,omitempty"`

	// Discount Discount `json:"discount,omitempty"`

	// If the subscription has ended, the date the subscription ended.
	EndedAt int `json:"ended_at,omitempty"`
	// Unique identifier for the object.
	Id            string                              `json:"id"`
	Items         []subscriptionItem.SubscriptionItem `json:"items"`
	LatestInvoice invoice.Invoice                     `json:"latest_invoice,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// Specifies the approximate timestamp on which any pending invoice items will be billed according to the schedule provided at `pending_invoice_item_interval`.
	NextPendingInvoiceItemInvoice int `json:"next_pending_invoice_item_invoice,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object                     string                                    `json:"object"`
	PauseCollection            SubscriptionsResourcePauseCollection      `json:"pause_collection,omitempty"`
	PaymentSettings            SubscriptionsResourcePaymentSettings      `json:"payment_settings,omitempty"`
	PendingInvoiceItemInterval SubscriptionPendingInvoiceItemInterval    `json:"pending_invoice_item_interval,omitempty"`
	PendingSetupIntent         intent.SetupIntent                        `json:"pending_setup_intent,omitempty"`
	PendingUpdate              SubscriptionsResourcePendingUpdate        `json:"pending_update,omitempty"`
	Schedule                   subscriptionSchedule.SubscriptionSchedule `json:"schedule,omitempty"`
	// Date when the subscription was first created. The date might differ from the `created` date due to backdating.
	StartDate int32 `json:"start_date"`
	// Possible values are `incomplete`, `incomplete_expired`, `trialing`, `active`, `past_due`, `canceled`, or `unpaid`.   For `collection_method=charge_automatically` a subscription moves into `incomplete` if the initial payment attempt fails. A subscription in this state can only have metadata and default_source updated. Once the first invoice is paid, the subscription moves into an `active` state. If the first invoice is not paid within 23 hours, the subscription transitions to `incomplete_expired`. This is a terminal state, the open invoice will be voided and no further invoices will be generated.   A subscription that is currently in a trial period is `trialing` and moves to `active` when the trial period is over.   If subscription `collection_method=charge_automatically` it becomes `past_due` when payment to renew it fails and `canceled` or `unpaid` (depending on your subscriptions settings) when Stripe has exhausted all payment retry attempts.   If subscription `collection_method=send_invoice` it becomes `past_due` when its invoice is not paid by the due date, and `canceled` or `unpaid` if it is still not paid by an additional deadline after that. Note that when a subscription has a status of `unpaid`, no subsequent invoices will be attempted (invoices will be created, but then immediately automatically closed). After receiving updated payment information from a customer, you may choose to reopen and pay their closed invoices.
	Status string `json:"status"`
	// If the subscription has a trial, the end of that trial.
	TrialEnd int `json:"trial_end,omitempty"`
	// If the subscription has a trial, the beginning of that trial.
	TrialStart int `json:"trial_start,omitempty"`
}
