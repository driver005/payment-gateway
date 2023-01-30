package methods

import "github.com/driver005/gateway/core"

type PaymentMethodIdeal struct {
	core.Model

	Bank string `json:"bank,omitempty"`
	Bic  string `json:"bic,omitempty"`
}
