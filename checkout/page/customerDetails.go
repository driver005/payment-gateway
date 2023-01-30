package page

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/address"
)

// PaymentPagesCheckoutSessionCustomerDetails
type PaymentPagesCheckoutSessionCustomerDetails struct {
	core.Model

	Address *address.Address `json:"address,omitempty" database:"foreignKey:id"`
	// The email associated with the Customer, if one exists, on the Checkout Session after a completed Checkout Session or at time of session expiry. Otherwise, if the customer has consented to promotional content, this value is the most recent valid email provided by the customer on the Checkout form.
	Email string `json:"email,omitempty"`
	// The customer's name after a completed Checkout Session. Note: This property is populated only for sessions on or after March 30, 2022.
	Name string `json:"name,omitempty"`
	// The customer's phone number after a completed Checkout Session.
	Phone string `json:"phone,omitempty"`
	// The customer’s tax exempt status after a completed Checkout Session.
	TaxExempt string `json:"tax_exempt,omitempty"`
	// The customer’s tax IDs after a completed Checkout Session.
	TaxIds []PaymentPagesCheckoutSessionTaxId `json:"tax_ids,omitempty" database:"foreignKey:id"`
}
