package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/payment/method"
)

// SetupAttemptPaymentMethodDetailsIdeal
type SetupAttemptPaymentMethodDetailsIdeal struct {
	core.Model

	// The customer's bank. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
	Bank string `json:"bank,omitempty"`
	// The Bank Identifier Code of the customer's bank.
	Bic                       string               `json:"bic,omitempty"`
	GeneratedSepaDebit        method.PaymentMethod `json:"generated_sepa_debi,omitempty" database:"foreignKey:id"`
	GeneratedSepaDebitMandate mandate.Mandate      `json:"generated_sepa_debit_mandate,omitempty" database:"foreignKey:id"`
	// Last four characters of the IBAN.
	IbanLast4 string `json:"iban_last4,omitempty"`
	// Owner's verified full name. Values are verified or provided by iDEAL directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name,omitempty"`
}
