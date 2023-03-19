package page

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// PaymentPagesCheckoutSessionShippingAddressCollection
type PaymentPagesCheckoutSessionShippingAddressCollection struct {
	core.Model

	// An array of two-letter ISO country codes representing which countries Checkout should provide as options for shipping locations. Unsupported country codes: `AS, CX, CC, CU, HM, IR, KP, MH, FM, NF, MP, PW, SD, SY, UM, VI`.
	AllowedCountries pq.StringArray `json:"allowed_countries,omitempty" database:"type:varchar(64)[]"`
}
