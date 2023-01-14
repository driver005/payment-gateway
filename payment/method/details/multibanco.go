package details

// PaymentMethodDetailsMultibanco
type PaymentMethodDetailsMultibanco struct {
	// Entity number associated with this Multibanco payment.
	Entity string `json:"entity,omitempty"`
	// Reference number associated with this Multibanco payment.
	Reference string `json:"reference,omitempty"`
}
