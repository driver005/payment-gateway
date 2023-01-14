package models

type AutomaticTax struct {
	// Whether Stripe automatically computes tax on this invoice. Note that incompatible invoice items (invoice items with manually specified [tax rates](https://stripe.com/docs/api/tax_rates), negative amounts, or `tax_behavior=unspecified`) cannot be added to automatic tax invoices.
	Enabled bool `json:"enabled"`
	// The status of the most recent automated tax calculation for this invoice.
	Status string `json:"status,omitempty"`
}