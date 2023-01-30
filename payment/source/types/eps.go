package types

import "github.com/driver005/gateway/core"

// SourceTypeEps struct for SourceTypeEps
type SourceTypeEps struct {
	core.Model

	Reference           string `json:"reference,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
