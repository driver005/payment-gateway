package fundingInstruction

// FundingInstructions Each customer has a [`balance`](https://stripe.com/docs/api/customers/object#customer_object-balance) that is automatically applied to future invoices and payments using the `customer_balance` payment method. Customers can fund this balance by initiating a bank transfer to any account in the `financial_addresses` field. Related guide: [Customer Balance - Funding Instructions](https://stripe.com/docs/payments/customer-balance/funding-instructions) to learn more
type FundingInstructions struct {
	BankTransfer FundingInstructionsBankTransfer `json:"bank_transfer"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// The `funding_type` of the returned instructions
	FundingType string `json:"funding_type"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
