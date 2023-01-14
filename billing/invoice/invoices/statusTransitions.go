package invoices

// InvoicesStatusTransitions
type InvoicesStatusTransitions struct {
	// The time that the invoice draft was finalized.
	FinalizedAt int `json:"finalized_at,omitempty"`
	// The time that the invoice was marked uncollectible.
	MarkedUncollectibleAt int `json:"marked_uncollectible_at,omitempty"`
	// The time that the invoice was paid.
	PaidAt int `json:"paid_at,omitempty"`
	// The time that the invoice was voided.
	VoidedAt int `json:"voided_at,omitempty"`
}
