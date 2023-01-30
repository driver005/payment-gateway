package methods

import "github.com/driver005/gateway/core"

type PaymentMethodCardPresent struct {
	core.Model

	Brand             string   `json:"brand,omitempty"`
	CardholderName    string   `json:"cardholder_name,omitempty"`
	Country           string   `json:"country,omitempty"`
	ExpMonth          int      `json:"exp_month,omitempty"`
	ExpYear           int      `json:"exp_year,omitempty"`
	Fingerprint       string   `json:"fingerprint,omitempty"`
	Funding           string   `json:"funding,omitempty"`
	Last4             string   `json:"last4,omitempty"`
	AvailableNetworks []string `json:"available_networks,omitempty" database:"type:text[]"`
	PreferredNetwork  string   `json:"preferred_network,omitempty"`
	ReadMethod        string   `json:"read_method,omitempty"`
}
