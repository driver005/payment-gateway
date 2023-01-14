package next

import "github.com/driver005/gateway/utils/fundingInstruction"

// PaymentIntentNextActionDisplayBankTransferInstructions
type PaymentIntentNextActionDisplayBankTransferInstructions struct {
	// The remaining amount that needs to be transferred to complete the payment.
	AmountRemaining int `json:"amount_remaining,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// A list of financial addresses that can be used to fund the customer balance
	FinancialAddresses []fundingInstruction.FundingInstructionsBankTransferFinancialAddress `json:"financial_addresses,omitempty"`
	// A link to a hosted page that guides your customer through completing the transfer.
	HostedInstructionsUrl string `json:"hosted_instructions_url,omitempty"`
	// A string identifying this payment. Instruct your customer to include this code in the reference or memo field of their bank transfer.
	Reference string `json:"reference,omitempty"`
	// Type of bank transfer
	Type string `json:"type"`
}
