package methods

// PaymentMethodCardWalletVisaCheckout
type PaymentMethodCardWalletVisaCheckout struct {
	BillingAddress PaymentMethodCardWalletMasterpassBillingAddress `json:"billing_address,omitempty"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email,omitempty"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name            string                                           `json:"name,omitempty"`
	ShippingAddress PaymentMethodCardWalletMasterpassShippingAddress `json:"shipping_address,omitempty"`
}
