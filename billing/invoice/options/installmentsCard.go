package options

import "github.com/driver005/gateway/core"

// InvoiceInstallmentsCard
type InvoiceInstallmentsCard struct {
	core.Model

	// Whether Installments are enabled for this Invoice.
	Enabled bool `json:"enabled,omitempty"`
}
