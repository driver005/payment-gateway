package options

import "github.com/driver005/gateway/core"

// CheckoutCardInstallmentsOptions
type CheckoutCardInstallmentsOptions struct {
	core.Model

	// Indicates if installments are enabled
	Enabled bool `json:"enabled,omitempty"`
}
