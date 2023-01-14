package options

import "github.com/driver005/gateway/billing/invoice/options"

// SubscriptionsResourcePaymentMethodOptionsUsBankAccount This sub-hash contains details about the ACH direct debit payment method options to pass to invoices created by the subscription.
type SubscriptionsResourcePaymentMethodOptionsUsBankAccount struct {
	InvoicePaymentMethodOptionsUsBankAccount *options.InvoicePaymentMethodOptionsUsBankAccount
}
