package types

import "github.com/driver005/gateway/core"

// SourceTypeAchDebit struct for SourceTypeAchDebit
type SourceTypeAchDebit struct {
	core.Model

	BankName      string `json:"bank_name,omitempty"`
	Country       string `json:"country,omitempty"`
	Fingerprint   string `json:"fingerprint,omitempty"`
	Last4         string `json:"last4,omitempty"`
	RoutingNumber string `json:"routing_number,omitempty"`
	Type          string `json:"type,omitempty"`
}
