package options

import "github.com/driver005/gateway/core"

// PaymentMethodOptionsWechatPay
type PaymentMethodOptionsWechatPay struct {
	core.Model

	// The app ID registered with WeChat Pay. Only required when client is ios or android.
	AppId string `json:"app_id,omitempty"`
	// The client type that the end customer will pay from
	Client string `json:"client,omitempty"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.  Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.  When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
}
