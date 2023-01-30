package methods

import (
	"github.com/driver005/gateway/core"
)

// LinkedAccountOptionsUsBankAccount
type LinkedAccountOptionsUsBankAccount struct {
	core.Model

	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []string `json:"permissions,omitempty" database:"type:text[]"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnUrl string `json:"return_url,omitempty"`
}

// PaymentIntentPaymentMethodOptionsUsBankAccount
type PaymentIntentPaymentMethodOptionsUsBankAccount struct {
	core.Model

	FinancialConnections *LinkedAccountOptionsUsBankAccount `json:"financial_connections,omitempty" database:"foreignKey:id"`
	// Indicates that you intend to make future payments with this PaymentIntent's payment method.  Providing this parameter will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete. If no Customer was provided, the payment method can still be [attached](https://stripe.com/docs/api/payment_methods/attach) to a Customer after the transaction completes.  When processing card payments, Stripe also uses `setup_future_usage` to dynamically optimize your payment flow and comply with regional legislation and network rules, such as [SCA](https://stripe.com/docs/strong-customer-authentication).
	SetupFutureUsage string `json:"setup_future_usage,omitempty"`
}
