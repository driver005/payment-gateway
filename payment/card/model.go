package card

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/models"
)

// Card You can store multiple cards on a customer in order to charge the customer later. You can also store multiple debit cards on a recipient in order to transfer to those cards later.  Related guide: [Card Payments with Sources](https://stripe.com/docs/sources/cards).
type Card struct {
	Account models.Account `json:"account,omitempty"`
	// City/District/Suburb/Town/Village.
	AddressCity string `json:"address_city,omitempty"`
	// Billing address country, if provided when creating card.
	AddressCountry string `json:"address_country,omitempty"`
	// Address line 1 (Street address/PO Box/Company name).
	AddressLine1 string `json:"address_line1,omitempty"`
	// If `address_line1` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check string `json:"address_line1_check,omitempty"`
	// Address line 2 (Apartment/Suite/Unit/Building).
	AddressLine2 string `json:"address_line2,omitempty"`
	// State/County/Province/Region.
	AddressState string `json:"address_state,omitempty"`
	// ZIP or postal code.
	AddressZip string `json:"address_zip,omitempty"`
	// If `address_zip` was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressZipCheck string `json:"address_zip_check,omitempty"`
	// A set of available payout methods for this card. Only values from this set should be passed as the `method` when creating a payout.
	AvailablePayoutMethods []string `json:"available_payout_methods,omitempty"`
	// Card brand. Can be `American Express`, `Diners Club`, `Discover`, `JCB`, `MasterCard`, `UnionPay`, `Visa`, or `Unknown`.
	Brand string `json:"brand"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country,omitempty"`
	// Three-letter [ISO code for currency](https://stripe.com/docs/payouts). Only applicable on accounts (not customers or recipients). The card can be used as a transfer destination for funds in this currency.
	Currency string `json:"currency,omitempty"`
	Customer customer.Customer `json:"customer,omitempty"`
	// If a CVC was provided, results of the check: `pass`, `fail`, `unavailable`, or `unchecked`. A result of unchecked indicates that CVC was provided but hasn't been checked yet. Checks are typically performed when attaching a card to a Customer object, or when creating a charge. For more details, see [Check if a card is valid without a charge](https://support.stripe.com/questions/check-if-a-card-is-valid-without-a-charge).
	CvcCheck string `json:"cvc_check,omitempty"`
	// Whether this card is the default external account for its currency.
	DefaultForCurrency bool `json:"default_for_currency,omitempty"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string `json:"dynamic_last4,omitempty"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who’ve signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.  *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata,omitempty"`
	// Cardholder name.
	Name string `json:"name,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// For external accounts, possible values are `new` and `errored`. If a transfer fails, the status is set to `errored` and transfers are stopped until account details are updated.
	Status string `json:"status,omitempty"`
	// If the card number is tokenized, this is the method that was used. Can be `android_pay` (includes Google Pay), `apple_pay`, `masterpass`, `visa_checkout`, or null.
	TokenizationMethod string `json:"tokenization_method,omitempty"`
}