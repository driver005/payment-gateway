package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Customer - Represents a customer
type Customer struct {
	core.Model

	// The customer's email
	Email string `json:"email"`

	// The customer's first name
	FirstName string `json:"first_name" database:"default:null"`

	// The customer's first name
	LastName string `json:"last_name" database:"default:null"`

	// The customer's billing address ID
	BillingAddressId uuid.NullUUID `json:"billing_address_id,omitempty"`

	BillingAddress *Address `json:"billing_address" database:"foreignKey:id;references:billing_address_id"`

	// Available if the relation `shipping_addresses` is expanded.
	ShippingAddresses []Address `json:"shipping_addresses" database:"foreignKey:id"`

	// The customer's phone number
	Phone string `json:"phone" database:"default:null"`

	// Whether the customer has an account or not
	HasAccount bool `json:"has_account" database:"default:null"`
}
