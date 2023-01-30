package settings

import "github.com/driver005/gateway/core"

// InvoiceSettingCustomField
type InvoiceSettingCustomField struct {
	core.Model

	// The name of the custom field.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}
