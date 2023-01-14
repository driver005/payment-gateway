package fundingInstruction

// FundingInstructionsBankTransferIbanRecord Iban Records contain E.U. bank account details per the SEPA format.
type FundingInstructionsBankTransferIbanRecord struct {
	// The name of the person or business that owns the bank account
	AccountHolderName string `json:"account_holder_name"`
	// The BIC/SWIFT code of the account.
	Bic string `json:"bic"`
	// Two-letter country code ([ISO 3166-1 alpha-2](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2)).
	Country string `json:"country"`
	// The IBAN of the account.
	Iban string `json:"iban"`
}
