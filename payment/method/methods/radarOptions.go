package methods

import "github.com/driver005/gateway/core"

type PaymentMethodRadarOptions struct {
	core.Model

	Session string `json:"session,omitempty"`
}
