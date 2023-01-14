package types

// SourceTypeAlipay struct for SourceTypeAlipay
type SourceTypeAlipay struct {
	DataString string `json:"data_string,omitempty"`
	NativeUrl string `json:"native_url,omitempty"`
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
}