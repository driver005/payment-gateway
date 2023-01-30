package types

import "github.com/driver005/gateway/core"

// SourceTypeWechat struct for SourceTypeWechat
type SourceTypeWechat struct {
	core.Model

	PrepayId            string `json:"prepay_id,omitempty"`
	QrCodeUrl           string `json:"qr_code_url,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}
