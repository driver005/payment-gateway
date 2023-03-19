package subscriptions

import (
	"github.com/driver005/gateway/billing/subscription/options"
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// SubscriptionsResourcePaymentSettings
type SubscriptionsResourcePaymentSettings struct {
	core.Model

	PaymentMethodOptions options.SubscriptionsResourcePaymentMethodOptions `json:"payment_method_options,omitempty" database:"foreignKey:id"`
	// The list of payment method types to provide to every invoice created by the subscription. If not set, Stripe attempts to automatically determine the types to use by looking at the invoice’s default payment method, the subscription’s default payment method, the customer’s default payment method, and your [invoice template settings](https://dashboard.stripe.com/settings/billing/invoice).
	PaymentMethodTypes pq.StringArray `json:"payment_method_types,omitempty" database:"type:varchar(64)[]"`
	// Either `off`, or `on_subscription`. With `on_subscription` Stripe updates `subscription.default_payment_method` when a subscription payment succeeds.
	SaveDefaultPaymentMethod string `json:"save_default_payment_method,omitempty"`
}

// SubscriptionsResourcePauseCollection The Pause Collection settings determine how we will pause collection for this subscription and for how long the subscription should be paused.
type SubscriptionsResourcePauseCollection struct {
	core.Model

	// The payment collection behavior for this subscription while paused. One of `keep_as_draft`, `mark_uncollectible`, or `void`.
	Behavior string `json:"behavior,omitempty"`
	// The time after which the subscription will resume collecting payments.
	ResumesAt int `json:"resumes_at,omitempty"`
}

// SubscriptionsResourcePendingUpdate Pending Updates store the changes pending from a previous update that will be applied to the Subscription upon successful payment.
type SubscriptionsResourcePendingUpdate struct {
	core.Model

	// If the update is applied, determines the date of the first full invoice, and, for plans with `month` or `year` intervals, the day of the month for subsequent invoices. The timestamp is in UTC format.
	BillingCycleAnchor int `json:"billing_cycle_anchor,omitempty"`
	// The point after which the changes reflected by this update will be discarded and no longer applied.
	ExpiresAt int `json:"expires_at,omitempty"`
	// Unix timestamp representing the end of the trial period the customer will get before being charged for the first time, if the update is applied.
	TrialEnd int `json:"trial_end,omitempty"`
	// Indicates if a plan's `trial_period_days` should be applied to the subscription. Setting `trial_end` per subscription is preferred, and this defaults to `false`. Setting this flag to `true` together with `trial_end` is not allowed. See [Using trial periods on subscriptions](https://stripe.com/docs/billing/subscriptions/trials) to learn more.
	TrialFromPlan bool `json:"trial_from_plan,omitempty"`
}
