package settings

import "github.com/driver005/gateway/core"

// InvoiceSettingRenderingOptions
type InvoiceSettingRenderingOptions struct {
	core.Model

	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display,omitempty"`
}
