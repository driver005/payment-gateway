package types

import "github.com/driver005/gateway/core"

// SourceTypeAlipay struct for SourceTypeAlipay
type SourceTypeAlipay struct {
	core.Model

	DataString          string `json:"data_string,omitempty"`
	NativeUrl           string `json:"native_url,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
