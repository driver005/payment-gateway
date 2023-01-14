package intent

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/payment/method"
)

// SetupIntentNextActionRedirectToUrl
type SetupIntentNextActionRedirectToUrl struct {
	// If the customer does not exit their browser while authenticating, they will be redirected to this specified URL after completion.
	ReturnUrl string `json:"return_url,omitempty"`
	// The URL you must redirect your customer to in order to authenticate.
	Url string `json:"url,omitempty"`
}

// SetupIntentNextActionVerifyWithMicrodeposits
type SetupIntentNextActionVerifyWithMicrodeposits struct {
	// The timestamp when the microdeposits are expected to land.
	ArrivalDate int `json:"arrival_date"`
	// The URL for the hosted verification page, which allows customers to verify their bank account.
	HostedVerificationUrl string `json:"hosted_verification_url"`
	// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
	MicrodepositType string `json:"microdeposit_type,omitempty"`
}

// SetupIntentNextAction
type SetupIntentNextAction struct {
	RedirectToUrl *SetupIntentNextActionRedirectToUrl `json:"redirect_to_url,omitempty"`
	// Type of the next action to perform, one of `redirect_to_url`, `use_stripe_sdk`, `alipay_handle_redirect`, `oxxo_display_details`, or `verify_with_microdeposits`.
	Type string `json:"type"`
	// When confirming a SetupIntent with Stripe.js, Stripe.js depends on the contents of this dictionary to invoke authentication flows. The shape of the contents is subject to change and is only intended to be used by Stripe.js.
	UseStripeSdk            map[string]interface{}                        `json:"use_stripe_sdk,omitempty"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits,omitempty"`
}


// SetupIntentTypeSpecificPaymentMethodOptionsClient
type SetupIntentTypeSpecificPaymentMethodOptionsClient struct {
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}


// SetupIntent A SetupIntent guides you through the process of setting up and saving a customer's payment credentials for future payments. For example, you could use a SetupIntent to set up and save your customer's card without immediately collecting a payment. Later, you can use [PaymentIntents](https://stripe.com/docs/api#payment_intents) to drive the payment flow.  Create a SetupIntent as soon as you're ready to collect your customer's payment credentials. Do not maintain long-lived, unconfirmed SetupIntents as they may no longer be valid. The SetupIntent then transitions through multiple [statuses](https://stripe.com/docs/payments/intents#intent-statuses) as it guides you through the setup process.  Successful SetupIntents result in payment credentials that are optimized for future payments. For example, cardholders in [certain regions](/guides/strong-customer-authentication) may need to be run through [Strong Customer Authentication](https://stripe.com/docs/strong-customer-authentication) at the time of payment method collection in order to streamline later [off-session payments](https://stripe.com/docs/payments/setup-intents). If the SetupIntent is used with a [Customer](https://stripe.com/docs/api#setup_intent_object-customer), upon success, it will automatically attach the resulting payment method to that Customer. We recommend using SetupIntents or [setup_future_usage](https://stripe.com/docs/api#payment_intent_object-setup_future_usage) on PaymentIntents to save payment methods in order to prevent saving invalid or unoptimized payment methods.  By using SetupIntents, you ensure that your customers experience the minimum set of required friction, even as regulations change over time.  Related guide: [Setup Intents API](https://stripe.com/docs/payments/setup-intents).
type SetupIntent struct {
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.  It can only be used for this Stripe Account’s own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `json:"attach_to_self,omitempty"`
	// Reason for cancellation of this SetupIntent, one of `abandoned`, `requested_by_customer`, or `duplicate`.
	CancellationReason string `json:"cancellation_reason,omitempty"`
	// The client secret of this SetupIntent. Used for client-side retrieval using a publishable key.  The client secret can be used to complete payment setup from your frontend. It should not be stored, logged, or exposed to anyone other than the customer. Make sure that you have TLS enabled on any page that includes the client secret.
	ClientSecret string `json:"client_secret,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created  int             `json:"created"`
	Customer customer.Customer `json:"customer,omitempty"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// Indicates the directions of money movement for which this payment method is intended to be used.  Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []string `json:"flow_directions,omitempty"`
	// Unique identifier for the object.
	Id             string           `json:"id"`
	LastSetupError models.ApiErrors `json:"last_setup_error,omitempty"`
	// LatestAttempt  *attempt.SetupAttempt `json:"latest_attempt,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool            `json:"livemode"`
	Mandate  mandate.Mandate `json:"mandate,omitempty"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata   map[string]string     `json:"metadata,omitempty"`
	NextAction SetupIntentNextAction `json:"next_action,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object        string               `json:"object"`
	PaymentMethod method.PaymentMethod `json:"payment_method,omitempty"`
	// PaymentMethodOptions  SetupIntentPaymentMethodOptions1 `json:"payment_method_options,omitempty"`
	// The list of payment method types (e.g. card) that this SetupIntent is allowed to set up.
	PaymentMethodTypes []string        `json:"payment_method_types"`
	SingleUseMandate   mandate.Mandate `json:"single_use_mandate,omitempty"`
	// [Status](https://stripe.com/docs/payments/intents#intent-statuses) of this SetupIntent, one of `requires_payment_method`, `requires_confirmation`, `requires_action`, `processing`, `canceled`, or `succeeded`.
	Status string `json:"status"`
	// Indicates how the payment method is intended to be used in the future.  Use `on_session` if you intend to only reuse the payment method when the customer is in your checkout flow. Use `off_session` if your customer may or may not be in your checkout flow. If not provided, this value defaults to `off_session`.
	Usage string `json:"usage"`
}
