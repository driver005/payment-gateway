package methods

// PaymentIntentPaymentMethodOptionsCard1 struct for PaymentIntentPaymentMethodOptionsCard1
type PaymentIntentPaymentMethodOptionsCard1 struct {
	PaymentIntentPaymentMethodOptionsCard

	// Controls when the funds will be captured from the customer's account.
	CaptureMethod string `json:"capture_method,omitempty"`
	// Installments  *paymentFlow.PaymentFlowsInstallmentOptions `json:"installments,omitempty"  database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
