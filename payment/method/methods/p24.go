package methods

import "github.com/driver005/gateway/core"

type PaymentMethodP24 struct {
	core.Model

	Bank string `json:"bank,omitempty"`
}
