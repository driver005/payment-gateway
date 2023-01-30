package options

import (
	"github.com/driver005/gateway/billing/invoice/options"
)

// SubscriptionsResourcePaymentMethodOptionsCustomerBalance This sub-hash contains details about the Bank transfer payment method options to pass to invoices created by the subscription.
type SubscriptionsResourcePaymentMethodOptionsCustomerBalance struct {
	options.InvoicePaymentMethodOptionsCustomerBalance
}
