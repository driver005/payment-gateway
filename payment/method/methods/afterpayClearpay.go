package methods

import "github.com/driver005/gateway/core"

type PaymentMethodAfterpayClearpay struct {
	core.Model

	Type string `json:"type,omitempty"`
}
