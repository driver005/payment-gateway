package methods

import "github.com/driver005/gateway/core"

type PaymentMethodEps struct {
	core.Model

	Bank string `json:"bank,omitempty"`
}
