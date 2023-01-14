package options

// PaymentMethodOptionsCustomerBalanceBankTransfer
type PaymentMethodOptionsCustomerBalanceBankTransfer struct {
	EuBankTransfer *PaymentMethodOptionsCustomerBalanceEuBankAccount `json:"eu_bank_transfer,omitempty"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.  Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes []string `json:"requested_address_types,omitempty"`
	// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type,omitempty"`
}
