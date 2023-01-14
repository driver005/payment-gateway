package settings

import "github.com/driver005/gateway/payment/method"

// InvoiceSettingCustomerSetting
type InvoiceSettingCustomerSetting struct {
	// Default custom fields to be displayed on invoices for this customer.
	CustomFields         []InvoiceSettingCustomField                       `json:"custom_fields,omitempty"`
	DefaultPaymentMethod method.PaymentMethod `json:"default_payment_method,omitempty"`
	// Default footer to be displayed on invoices for this customer.
	Footer           string                                        `json:"footer,omitempty"`
	RenderingOptions InvoiceSettingCustomerSettingRenderingOptions `json:"rendering_options,omitempty"`
}
