package types

import "github.com/driver005/gateway/core"

// SourceTypeAchCreditTransfer struct for SourceTypeAchCreditTransfer
type SourceTypeAchCreditTransfer struct {
	core.Model

	AccountNumber           string `json:"account_number,omitempty"`
	BankName                string `json:"bank_name,omitempty"`
	Fingerprint             string `json:"fingerprint,omitempty"`
	RefundAccountHolderName string `json:"refund_account_holder_name,omitempty"`
	RefundAccountHolderType string `json:"refund_account_holder_type,omitempty"`
	RefundRoutingNumber     string `json:"refund_routing_number,omitempty"`
	RoutingNumber           string `json:"routing_number,omitempty"`
	SwiftCode               string `json:"swift_code,omitempty"`
}
