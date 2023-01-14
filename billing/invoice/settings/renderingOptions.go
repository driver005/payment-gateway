package settings

// InvoiceSettingRenderingOptions
type InvoiceSettingRenderingOptions struct {
	// How line-item prices and amounts will be displayed with respect to tax on invoice PDFs.
	AmountTaxDisplay string `json:"amount_tax_display,omitempty"`
}
