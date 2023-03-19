package options

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions
type InvoicePaymentMethodOptionsUsBankAccountLinkedAccountOptions struct {
	core.Model

	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions pq.StringArray `json:"permissions,omitempty" database:"type:varchar(64)[]"`
}
