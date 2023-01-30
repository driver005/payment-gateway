package options

import "github.com/driver005/gateway/core"

// PaymentLinksResourceShippingAddressCollection
type PaymentLinksResourceShippingAddressCollection struct {
	core.Model

	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries []string `json:"allowed_countries,omitempty"`
}
