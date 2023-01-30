package methods

import (
	"github.com/driver005/gateway/core"
)

type PaymentMethodCard struct {
	core.Model

	Brand                      string                  `json:"brand,omitempty"`
	AddressLine1Check          string                  `json:"address_line1_check,omitempty"`
	AddressPostalCodeCheck     string                  `json:"address_postal_code_check,omitempty"`
	CvcCheck                   string                  `json:"cvc_check,omitempty"`
	Country                    string                  `json:"country,omitempty"`
	ExpMonth                   int                     `json:"exp_month,omitempty"`
	ExpYear                    int                     `json:"exp_year,omitempty"`
	Fingerprint                string                  `json:"fingerprint,omitempty"`
	Funding                    string                  `json:"funding,omitempty"`
	Last4                      string                  `json:"last4,omitempty"`
	AvailableNetworks          []string                `json:"available_networks,omitempty" database:"type:text[]"`
	PreferredNetwork           string                  `json:"preferred_network,omitempty"`
	ThreeDSecureUsageSupported bool                    `json:"three_d_secure_usage_supported,omitempty"`
	Wallet                     PaymentMethodCardWallet `json:"wallet,omitempty" database:"foreignKey:id"`
}
