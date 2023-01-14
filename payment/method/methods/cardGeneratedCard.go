package methods

// PaymentMethodCardGeneratedCard
type PaymentMethodCardGeneratedCard struct {
	// The charge that created this object.
	Charge               string                                             `json:"charge,omitempty"`
	PaymentMethodDetails PaymentMethodCardGeneratedCardPaymentMethodDetails `json:"payment_method_details,omitempty"`
	SetupAttempt         PaymentMethodCardGeneratedCardSetupAttempt         `json:"setup_attempt,omitempty"`
}
