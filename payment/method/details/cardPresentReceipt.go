package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsCardPresentReceipt
type PaymentMethodDetailsCardPresentReceipt struct {
	core.Model

	// The type of account being debited or credited
	AccountType string `json:"account_type,omitempty"`
	// EMV tag 9F26, cryptogram generated by the integrated circuit chip.
	ApplicationCryptogram string `json:"application_cryptogram,omitempty"`
	// Mnenomic of the Application Identifier.
	ApplicationPreferredName string `json:"application_preferred_name,omitempty"`
	// Identifier for this transaction.
	AuthorizationCode string `json:"authorization_code,omitempty"`
	// EMV tag 8A. A code returned by the card issuer.
	AuthorizationResponseCode string `json:"authorization_response_code,omitempty"`
	// How the cardholder verified ownership of the card.
	CardholderVerificationMethod string `json:"cardholder_verification_method,omitempty"`
	// EMV tag 84. Similar to the application identifier stored on the integrated circuit chip.
	DedicatedFileName string `json:"dedicated_file_name,omitempty"`
	// The outcome of a series of EMV functions performed by the card reader.
	TerminalVerificationResults string `json:"terminal_verification_results,omitempty"`
	// An indication of various EMV functions performed during the transaction.
	TransactionStatusInformation string `json:"transaction_status_information,omitempty"`
}