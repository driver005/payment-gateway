package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method/options"
)

// PaymentIntentPaymentMethodOptionsCard
type PaymentIntentPaymentMethodOptionsCard struct {
	core.Model

	Installments   *options.PaymentMethodOptionsCardInstallments   `json:"installments,omitempty" database:"foreignKey:id"`
	MandateOptions *options.PaymentMethodOptionsCardMandateOptions `json:"mandate_options,omitempty"  database:"foreignKey:id"`
	// Selected network to process this payment intent on. Depends on the available networks of the card attached to the payment intent. Can be only set confirm-time.
	Network string `json:"network,omitempty"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Permitted values include: `automatic` or `any`. If not provided, defaults to `automatic`. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure string `json:"request_three_d_secure,omitempty"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.  Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.  When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kana prefix (shortened Kana descriptor) or Kana statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 22 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 22 characters.
	StatementDescriptorSuffixKana string `json:"statement_descriptor_suffix_kana,omitempty"`
	// Provides information about a card payment that customers see on their statements. Concatenated with the Kanji prefix (shortened Kanji descriptor) or Kanji statement descriptor that’s set on the account to form the complete statement descriptor. Maximum 17 characters. On card statements, the *concatenation* of both prefix and suffix (including separators) will appear truncated to 17 characters.
	StatementDescriptorSuffixKanji string `json:"statement_descriptor_suffix_kanji,omitempty"`
}
