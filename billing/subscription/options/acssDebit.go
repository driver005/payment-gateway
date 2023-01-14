package options

import "github.com/driver005/gateway/billing/invoice/options"

// SubscriptionsResourcePaymentMethodOptionsAcssDebit This sub-hash contains details about the Canadian pre-authorized debit payment method options to pass to invoices created by the subscription.
type SubscriptionsResourcePaymentMethodOptionsAcssDebit struct {
	InvoicePaymentMethodOptionsAcssDebit *options.InvoicePaymentMethodOptionsAcssDebit
}
