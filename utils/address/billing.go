package address

import "github.com/driver005/gateway/core"

// BillingDetails
type BillingDetails struct {
	core.Model

	Address *Address `json:"address,omitempty" database:"foreignKey:id"`
	// Email address.
	Email string `json:"email,omitempty"`
	// Full name.
	Name string `json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone string `json:"phone,omitempty"`
}
