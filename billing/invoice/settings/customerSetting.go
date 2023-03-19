package settings

import "github.com/driver005/gateway/core"

// InvoiceSettingCustomerSetting
type InvoiceSettingCustomerSetting struct {
	core.Model

	// Default custom fields to be displayed on invoices for this customer.
	CustomFields []InvoiceSettingCustomField `json:"custom_fields,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// DefaultPaymentMethod method.PaymentMethod `json:"default_payment_method,omitempty"`
	// Default footer to be displayed on invoices for this customer.
	Footer           string                         `json:"footer,omitempty"`
	RenderingOptions InvoiceSettingRenderingOptions `json:"rendering_options,omitempty" database:"foreignKey:id"`
}
