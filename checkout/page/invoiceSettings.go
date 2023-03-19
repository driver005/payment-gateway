package page

import (
	"github.com/driver005/gateway/billing/invoice/settings"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/tax"
)

// PaymentPagesCheckoutSessionInvoiceSettings
type PaymentPagesCheckoutSessionInvoiceSettings struct {
	core.Model

	// The account tax IDs associated with the invoice.
	AccountTaxIds []tax.TaxId `json:"account_tax_ids,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// Custom fields displayed on the invoice.
	CustomFields []settings.InvoiceSettingCustomField `json:"custom_fields,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description,omitempty"`
	// Footer displayed on the invoice.
	Footer           string                                   `json:"footer,omitempty"`
	RenderingOptions *settings.InvoiceSettingRenderingOptions `json:"rendering_options,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
