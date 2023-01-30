package methods

import "github.com/driver005/gateway/core"

type PaymentMethodSofort struct {
	core.Model

	Country string `json:"country,omitempty"`
}
