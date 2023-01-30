package options

import "github.com/driver005/gateway/core"

// InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions
type InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions struct {
	core.Model

	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []string `json:"permissions,omitempty" database:"type:text[]"`
}
