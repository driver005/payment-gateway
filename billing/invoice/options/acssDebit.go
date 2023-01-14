package options

// InvoicePaymentMethodOptionsAcssDebit
type InvoicePaymentMethodOptionsAcssDebit struct {
	MandateOptions *InvoicePaymentMethodOptionsAcssDebitMandateOptions `json:"mandate_options,omitempty"`
	// Bank account verification method.
	VerificationMethod string `json:"verification_method,omitempty"`
}
