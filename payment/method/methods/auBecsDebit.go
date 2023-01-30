package methods

import "github.com/driver005/gateway/core"

type PaymentMethodAuBecsDebit struct {
	core.Model

	BsbNumber   string `json:"bsb_number,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4       string `json:"last4,omitempty"`
}
