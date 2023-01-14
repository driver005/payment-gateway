package types

// SourceTypeAuBecsDebit struct for SourceTypeAuBecsDebit
type SourceTypeAuBecsDebit struct {
	BsbNumber string `json:"bsb_number,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4 string `json:"last4,omitempty"`
}
