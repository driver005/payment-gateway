package fundingInstruction

import "github.com/driver005/gateway/core"

// FundingInstructionsBankTransferIbanRecord Iban Records contain E.U. bank account details per the SEPA format.
type FundingInstructionsBankTransferIbanRecord struct {
	core.Model
	// The name of the person or business that owns the bank account
	AccountHolderName string `json:"account_holder_name,omitempty"`
	// The BIC/SWIFT code of the account.
	Bic string `json:"bic,omitempty"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country,omitempty"`
	// The IBAN of the account.
	Iban string `json:"iban,omitempty"`
}
