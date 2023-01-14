package details

// PaymentMethodDetailsCardInstallmentsPlan
type PaymentMethodDetailsCardInstallmentsPlan struct {
	// For `fixed_count` installment plans, this is the number of installment payments your customer will make to their credit card.
	Count int `json:"count,omitempty"`
	// For `fixed_count` installment plans, this is the interval between installment payments your customer will make to their credit card. One of `month`.
	Interval string `json:"interval,omitempty"`
	// Type of installment plan, one of `fixed_count`.
	Type string `json:"type"`
}
