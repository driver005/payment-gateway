package types

import "github.com/driver005/gateway/core"

// SourceTypeAuBecsDebit struct for SourceTypeAuBecsDebit
type SourceTypeAuBecsDebit struct {
	core.Model

	BsbNumber   string `json:"bsb_number,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4       string `json:"last4,omitempty"`
}
