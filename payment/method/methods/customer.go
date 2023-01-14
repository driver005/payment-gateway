package methods

import "github.com/driver005/gateway/models"

// PaymentMethodCustomer The ID of the Customer to which this PaymentMethod is saved. This will not be set when the PaymentMethod has not been saved to a Customer.
type PaymentMethodCustomer struct {
	Customer *models.Customer
	string   string
}
