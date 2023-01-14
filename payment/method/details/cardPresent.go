package details

// PaymentMethodDetailsCardPresent
type PaymentMethodDetailsCardPresent struct {
	// The authorized amount
	AmountAuthorized int `json:"amount_authorized,omitempty"`
	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand string `json:"brand,omitempty"`
	// When using manual capture, a future timestamp after which the charge will be automatically refunded if uncaptured.
	CaptureBefore *int `json:"capture_before,omitempty"`
	// The cardholder name as read from the card, in [ISO 7813](https://en.wikipedia.org/wiki/ISO/IEC_7813) format. May include alphanumeric characters, special characters and first/last name separator (`/`). In some cases, the cardholder name may not be available depending on how the issuer has configured the card. Cardholder name is typically not available on swipe or contactless payments, such as those made with Apple Pay and Google Pay.
	CardholderName string `json:"cardholder_name,omitempty"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country,omitempty"`
	// Authorization response cryptogram.
	EmvAuthData string `json:"emv_auth_data,omitempty"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.  *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding,omitempty"`
	// ID of a card PaymentMethod generated from the card_present PaymentMethod that may be attached to a Customer for future transactions. Only present if it was possible to generate a card PaymentMethod.
	GeneratedCard string `json:"generated_card,omitempty"`
	// Whether this [PaymentIntent](https://stripe.com/docs/api/payment_intents) is eligible for incremental authorizations. Request support using [request_incremental_authorization_support](https://stripe.com/docs/api/payment_intents/create#create_payment_intent-payment_method_options-card_present-request_incremental_authorization_support).
	IncrementalAuthorizationSupported bool `json:"incremental_authorization_supported"`
	// The last four digits of the card.
	Last4 string `json:"last4,omitempty"`
	// Identifies which network this charge was processed on. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `interac`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Network string `json:"network,omitempty"`
	// Defines whether the authorized amount can be over-captured or not
	OvercaptureSupported bool `json:"overcapture_supported"`
	// How card details were read in this transaction.
	ReadMethod string                                  `json:"read_method,omitempty"`
	Receipt    PaymentMethodDetailsCardPresentReceipt1 `json:"receipt,omitempty"`
}
