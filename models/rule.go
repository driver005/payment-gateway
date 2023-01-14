package models

// Rule
type Rule struct {
	// The action taken on the payment.
	Action string `json:"action"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// The predicate to evaluate the payment against.
	Predicate string `json:"predicate"`
}
