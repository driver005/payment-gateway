package source

import (
	"github.com/driver005/gateway/payment/source/types"
	"github.com/driver005/gateway/utils/address"
)

// SourceRedirectFlow
type SourceRedirectFlow struct {
	// The failure reason for the redirect, either `user_abort` (the customer aborted or dropped out of the redirect flow), `declined` (the authentication failed or the transaction was declined), or `processing_error` (the redirect failed due to a technical error). Present only if the redirect status is `failed`.
	FailureReason string `json:"failure_reason,omitempty"`
	// The URL you provide to redirect the customer to after they authenticated their payment.
	ReturnUrl string `json:"return_url"`
	// The status of the redirect, either `pending` (ready to be used by your customer to authenticate the transaction), `succeeded` (succesful authentication, cannot be reused) or `not_required` (redirect should not be used) or `failed` (failed authentication, cannot be reused).
	Status string `json:"status"`
	// The URL provided to you to redirect a customer to as part of a `redirect` authentication flow.
	Url string `json:"url"`
}

// SourceReceiverFlow
type SourceReceiverFlow struct {
	// The address of the receiver source. This is the value that should be communicated to the customer to send their funds to.
	Address string `json:"address,omitempty"`
	// The total amount that was moved to your balance. This is almost always equal to the amount charged. In rare cases when customers deposit excess funds and we are unable to refund those, those funds get moved to your balance and show up in amount_charged as well. The amount charged is expressed in the source's currency.
	AmountCharged int `json:"amount_charged"`
	// The total amount received by the receiver source. `amount_received = amount_returned + amount_charged` should be true for consumed sources unless customers deposit excess funds. The amount received is expressed in the source's currency.
	AmountReceived int `json:"amount_received"`
	// The total amount that was returned to the customer. The amount returned is expressed in the source's currency.
	AmountReturned int `json:"amount_returned"`
	// Type of refund attribute method, one of `email`, `manual`, or `none`.
	RefundAttributesMethod string `json:"refund_attributes_method"`
	// Type of refund attribute status, one of `missing`, `requested`, or `available`.
	RefundAttributesStatus string `json:"refund_attributes_status"`
}

// SourceOwner
type SourceOwner struct {
	Address address.Address `json:"address,omitempty"`
	// Owner's email address.
	Email string `json:"email,omitempty"`
	// Owner's full name.
	Name string `json:"name,omitempty"`
	// Owner's phone number (including extension).
	Phone           string          `json:"phone,omitempty"`
	VerifiedAddress address.Address `json:"verified_address,omitempty"`
	// Verified owner's email address. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedEmail string `json:"verified_email,omitempty"`
	// Verified owner's full name. Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedName string `json:"verified_name,omitempty"`
	// Verified owner's phone number (including extension). Verified values are verified or provided by the payment method directly (and if supported) at the time of authorization or settlement. They cannot be set or mutated.
	VerifiedPhone string `json:"verified_phone,omitempty"`
}

// SourceOrderItem
type SourceOrderItem struct {
	// The amount (price) for this order item.
	Amount int `json:"amount,omitempty"`
	// This currency of this order item. Required when `amount` is present.
	Currency string `json:"currency,omitempty"`
	// Human-readable description for this order item.
	Description string `json:"description,omitempty"`
	// The ID of the associated object for this line item. Expandable if not null (e.g., expandable to a SKU).
	Parent string `json:"parent,omitempty"`
	// The quantity of this order item. When type is `sku`, this is the number of instances of the SKU to be ordered.
	Quantity *int `json:"quantity,omitempty"`
	// The type of this order item. Must be `sku`, `tax`, or `shipping`.
	Type string `json:"type,omitempty"`
}

// SourceOrder
type SourceOrder struct {
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount for the order.
	Amount int `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// The email address of the customer placing the order.
	Email *string `json:"email,omitempty"`
	// List of items constituting the order.
	Items    []SourceOrderItem `json:"items,omitempty"`
	Shipping *address.Shipping `json:"shipping,omitempty"`
}

// SourceCodeVerificationFlow
type SourceCodeVerificationFlow struct {
	// The number of attempts remaining to authenticate the source object with a verification code.
	AttemptsRemaining int `json:"attempts_remaining"`
	// The status of the code verification, either `pending` (awaiting verification, `attempts_remaining` should be greater than 0), `succeeded` (successful verification) or `failed` (failed verification, cannot be verified anymore as `attempts_remaining` should be 0).
	Status string `json:"status"`
}

// Source `Source` objects allow you to accept a variety of payment methods. They represent a customer's payment instrument, and can be used with the Stripe API just like a `Card` object: once chargeable, they can be charged, or can be attached to customers.  Stripe doesn't recommend using the deprecated [Sources API](https://stripe.com/docs/api/sources). We recommend that you adopt the [PaymentMethods API](https://stripe.com/docs/api/payment_methods). This newer API provides access to our latest features and payment method types.  Related guides: [Sources API](https://stripe.com/docs/sources) and [Sources & Customers](https://stripe.com/docs/sources/customers).
type Source struct {
	AchCreditTransfer *types.SourceTypeAchCreditTransfer `json:"ach_credit_transfer,omitempty"`
	AchDebit          *types.SourceTypeAchDebit          `json:"ach_debit,omitempty"`
	AcssDebit         *types.SourceTypeAcssDebit         `json:"acss_debit,omitempty"`
	Alipay            *types.SourceTypeAlipay            `json:"alipay,omitempty"`
	// A positive integer in the smallest currency unit (that is, 100 cents for $1.00, or 1 for ¥1, Japanese Yen being a zero-decimal currency) representing the total amount associated with the source. This is the amount for which the source will be chargeable once ready. Required for `single_use` sources.
	Amount      int                          `json:"amount,omitempty"`
	AuBecsDebit *types.SourceTypeAuBecsDebit `json:"au_becs_debit,omitempty"`
	Bancontact  *types.SourceTypeBancontact  `json:"bancontact,omitempty"`
	Card        *types.SourceTypeCard        `json:"card,omitempty"`
	CardPresent *types.SourceTypeCardPresent `json:"card_present,omitempty"`
	// The client secret of the source. Used for client-side retrieval using a publishable key.
	ClientSecret     string                      `json:"client_secret"`
	CodeVerification *SourceCodeVerificationFlow `json:"code_verification,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) associated with the source. This is the currency for which the source will be chargeable once ready. Required for `single_use` sources.
	Currency string `json:"currency,omitempty"`
	// The ID of the customer to which this source is attached. This will not be present when the source has not been attached to a customer.
	Customer *string              `json:"customer,omitempty"`
	Eps      *types.SourceTypeEps `json:"eps,omitempty"`
	// The authentication `flow` of the source. `flow` is one of `redirect`, `receiver`, `code_verification`, `none`.
	Flow    string                   `json:"flow"`
	Giropay *types.SourceTypeGiropay `json:"giropay,omitempty"`
	// Unique identifier for the object.
	Id     string                  `json:"id"`
	Ideal  *types.SourceTypeIdeal  `json:"ideal,omitempty"`
	Klarna *types.SourceTypeKlarna `json:"klarna,omitempty"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata   map[string]string           `json:"metadata,omitempty"`
	Multibanco *types.SourceTypeMultibanco `json:"multibanco,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object      string                     `json:"object"`
	Owner       SourceOwner                `json:"owner,omitempty"`
	P24         *types.SourceTypeP24       `json:"p24,omitempty"`
	Receiver    *SourceReceiverFlow        `json:"receiver,omitempty"`
	Redirect    *SourceRedirectFlow        `json:"redirect,omitempty"`
	SepaDebit   *types.SourceTypeSepaDebit `json:"sepa_debit,omitempty"`
	Sofort      *types.SourceTypeSofort    `json:"sofort,omitempty"`
	SourceOrder *SourceOrder               `json:"source_order,omitempty"`
	// Extra information about a source. This will appear on your customer's statement every time you charge the source.
	StatementDescriptor string `json:"statement_descriptor,omitempty"`
	// The status of the source, one of `canceled`, `chargeable`, `consumed`, `failed`, or `pending`. Only `chargeable` sources can be used to create a charge.
	Status       string                        `json:"status"`
	ThreeDSecure *types.SourceTypeThreeDSecure `json:"three_d_secure,omitempty"`
	// The `type` of the source. The `type` is a payment method, one of `ach_credit_transfer`, `ach_debit`, `alipay`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `multibanco`, `klarna`, `p24`, `sepa_debit`, `sofort`, `three_d_secure`, or `wechat`. An additional hash is included on the source with a name matching this value. It contains additional information specific to the [payment method](https://stripe.com/docs/sources) used.
	Type string `json:"type"`
	// Either `reusable` or `single_use`. Whether this source should be reusable or not. Some source types may or may not be reusable by construction, while others may leave the option at creation. If an incompatible value is passed, an error will be returned.
	Usage  string                  `json:"usage,omitempty"`
	Wechat *types.SourceTypeWechat `json:"wechat,omitempty"`
}
