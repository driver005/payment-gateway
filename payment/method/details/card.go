package details

import "github.com/driver005/gateway/core"

type ThreeDSecureDetails struct {
	core.Model

	// For authenticated transactions: how the customer was authenticated by the issuing bank.
	AuthenticationFlow string `json:"authentication_flow,omitempty"`
	// Indicates the outcome of 3D Secure authentication.
	Result string `json:"result,omitempty"`
	// Additional information about why 3D Secure succeeded or failed based on the `result`.
	ResultReason string `json:"result_reason,omitempty"`
	// The version of 3D Secure that was used.
	Version string `json:"version,omitempty"`
}

// PaymentMethodDetailsCard
type PaymentMethodDetailsCard struct {
	core.Model

	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand string `json:"brand,omitempty"`
	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check string `json:"address_line1_check,omitempty"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck string `json:"address_postal_code_check,omitempty"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CvcCheck string `json:"cvc_check,omitempty"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country,omitempty"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.  *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding      string                               `json:"funding,omitempty"`
	Installments PaymentMethodDetailsCardInstallments `json:"installments,omitempty" database:"foreignKey:id"`
	// The last four digits of the card.
	Last4 string `json:"last4,omitempty"`
	// ID of the mandate used to make this payment or created by it.
	Mandate string `json:"mandate,omitempty"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network      string                         `json:"network,omitempty"`
	ThreeDSecure ThreeDSecureDetails            `json:"three_d_secure,omitempty" database:"foreignKey:id"`
	Wallet       PaymentMethodDetailsCardWallet `json:"wallet,omitempty" database:"foreignKey:id"`
}
