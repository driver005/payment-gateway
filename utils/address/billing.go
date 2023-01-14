package address

// BillingDetails
type BillingDetails struct {
	Address Address `json:"address,omitempty"`
	// Email address.
	Email string `json:"email,omitempty"`
	// Full name.
	Name string `json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone string `json:"phone,omitempty"`
}
