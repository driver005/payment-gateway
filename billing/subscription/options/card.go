package options

import (
	"github.com/driver005/gateway/core"
)

// SubscriptionsResourcePaymentMethodOptionsCard This sub-hash contains details about the Card payment method options to pass to invoices created by the subscription.
type SubscriptionsResourcePaymentMethodOptionsCard struct {
	core.Model

	// Amount to be charged for future payments.
	Amount int `json:"amount,omitempty"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType string `json:"amount_type,omitempty"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
	// Selected network to process this Subscription on. Depends on the available networks of the card attached to the Subscription. Can be only set confirm-time.
	Network string `json:"network,omitempty"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure string `json:"request_three_d_secure,omitempty"`
}
