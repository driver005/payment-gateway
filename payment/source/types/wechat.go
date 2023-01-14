package types

// SourceTypeWechat struct for SourceTypeWechat
type SourceTypeWechat struct {
	PrepayId *string `json:"prepay_id,omitempty"`
	QrCodeUrl string `json:"qr_code_url,omitempty"`
	StatementDescriptor *string `json:"statement_descriptor,omitempty"`
}
