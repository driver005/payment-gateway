package details

// PaymentMethodDetailsCard
type PaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand  string                          `json:"brand,omitempty"`
	Checks PaymentMethodDetailsCardChecks1 `json:"checks,omitempty"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country,omitempty"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.  *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding      string                                `json:"funding,omitempty"`
	Installments PaymentMethodDetailsCardInstallments1 `json:"installments,omitempty"`
	// The last four digits of the card.
	Last4 string `json:"last4,omitempty"`
	// ID of the mandate used to make this payment or created by it.
	Mandate string `json:"mandate,omitempty"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network      string                               `json:"network,omitempty"`
	ThreeDSecure PaymentMethodDetailsCardThreeDSecure `json:"three_d_secure,omitempty"`
	Wallet       PaymentMethodDetailsCardWallet1      `json:"wallet,omitempty"`
}
