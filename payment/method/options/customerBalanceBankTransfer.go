package options

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

// PaymentMethodOptionsCustomerBalanceBankTransfer
type PaymentMethodOptionsCustomerBalanceBankTransfer struct {
	core.Model

	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	EuBankTransferCountry string `json:"eu_bank_transfer_country,omitempty"`
	// List of address types that should be returned in the financial_addresses response. If not specified, all valid types will be returned.  Permitted values include: `sort_code`, `zengin`, `iban`, or `spei`.
	RequestedAddressTypes pq.StringArray `json:"requested_address_types,omitempty" database:"type:varchar(64)[]"`
	// The bank transfer type that this PaymentIntent is allowed to use for funding Permitted values include: `eu_bank_transfer`, `gb_bank_transfer`, `jp_bank_transfer`, or `mx_bank_transfer`.
	Type string `json:"type,omitempty"`
}

// PaymentMethodOptionsCustomerBalanceEuBankAccount
type PaymentMethodOptionsCustomerBalanceEuBankAccount struct {
	core.Model

	// The desired country code of the bank account information. Permitted values include: `DE`, `ES`, `FR`, `IE`, or `NL`.
	Country string `json:"country,omitempty"`
}
