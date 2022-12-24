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
	BillingAddressId uuid.NullUUID `json:"billing_address_id" database:"default:null"`

	BillingAddress *Address `json:"billing_address" database:"foreignKey:id;references:billing_address_id"`

	// Available if the relation `shipping_addresses` is expanded.
	ShippingAddresses []Address `json:"shipping_addresses" database:"foreignKey:id"`

	// The customer's phone number
	Phone string `json:"phone" database:"default:null"`

	// Whether the customer has an account or not
	HasAccount bool `json:"has_account" database:"default:null"`

	// Available if the relation `orders` is expanded.
	Orders []Order `json:"orders" database:"foreignKey:id"`

	// The customer groups the customer belongs to. Available if the relation `groups` is expanded.
	Groups []CustomerGroup `json:"groups" database:"many2many:customer_group"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
