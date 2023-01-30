package methods

import "github.com/driver005/gateway/core"

type PaymentMethodBacsDebit struct {
	core.Model

	Fingerprint string `json:"fingerprint,omitempty"`
	Last4       string `json:"last4,omitempty"`
	SortCode    string `json:"sort_code,omitempty"`
}
