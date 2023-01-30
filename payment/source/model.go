package source

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/source/types"
	"github.com/driver005/gateway/utils/address"
	"github.com/google/uuid"
)

// SourceRedirectFlow
type SourceRedirectFlow struct {
	core.Model

	FailureReason string `json:"failure_reason,omitempty"`
	ReturnUrl     string `json:"return_url,omitempty"`
	Status        string `json:"status,omitempty"`
	Url           string `json:"url,omitempty"`
}

// SourceReceiverFlow
type SourceReceiverFlow struct {
	core.Model

	Address                string `json:"address,omitempty"`
	AmountCharged          int    `json:"amount_charged,omitempty"`
	AmountReceived         int    `json:"amount_received,omitempty"`
	AmountReturned         int    `json:"amount_returned,omitempty"`
	RefundAttributesMethod string `json:"refund_attributes_method,omitempty"`
	RefundAttributesStatus string `json:"refund_attributes_status,omitempty"`
}

// SourceOwner
type SourceOwner struct {
	core.Model

	Address         *address.Address `json:"address,omitempty" database:"foreignKey:id"`
	Email           string           `json:"email,omitempty"`
	Name            string           `json:"name,omitempty"`
	Phone           string           `json:"phone,omitempty"`
	VerifiedAddress *address.Address `json:"verified_address,omitempty" database:"foreignKey:id"`
	VerifiedEmail   string           `json:"verified_email,omitempty"`
	VerifiedName    string           `json:"verified_name,omitempty"`
	VerifiedPhone   string           `json:"verified_phone,omitempty"`
}

// SourceOrderItem
type SourceOrderItem struct {
	core.Model

	Amount      int    `json:"amount,omitempty"`
	Currency    string `json:"currency,omitempty"`
	Description string `json:"description,omitempty"`
	Parent      string `json:"parent,omitempty"`
	Quantity    int    `json:"quantity,omitempty"`
	Type        string `json:"type,omitempty"`
}

// SourceOrder
type SourceOrder struct {
	core.Model

	Amount   int               `json:"amount,omitempty"`
	Currency string            `json:"currency,omitempty"`
	Email    string            `json:"email,omitempty"`
	Items    []SourceOrderItem `json:"items,omitempty" database:"foreignKey:id"`
	Shipping *address.Shipping `json:"shipping,omitempty" database:"foreignKey:id"`
}

// SourceCodeVerificationFlow
type SourceCodeVerificationFlow struct {
	core.Model

	AttemptsRemaining int    `json:"attempts_remaining,omitempty"`
	Status            string `json:"status,omitempty"`
}

// Source `Source` objects allow you to accept a variety of payment methods. They represent a customer's payment instrument, and can be used with the Stripe API just like a `Card` object: once chargeable, they can be charged, or can be attached to customers.  Stripe doesn't recommend using the deprecated [Sources API](https://stripe.com/docs/api/sources). We recommend that you adopt the [PaymentMethods API](https://stripe.com/docs/api/payment_methods). This newer API provides access to our latest features and payment method types.  Related guides: [Sources API](https://stripe.com/docs/sources) and [Sources & Customers](https://stripe.com/docs/sources/customers).
type Source struct {
	core.Model

	AchCreditTransfer   *types.SourceTypeAchCreditTransfer `json:"ach_credit_transfer,omitempty" database:"foreignKey:id"`
	AchDebit            *types.SourceTypeAchDebit          `json:"ach_debit,omitempty" database:"foreignKey:id"`
	AcssDebit           *types.SourceTypeAcssDebit         `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Alipay              *types.SourceTypeAlipay            `json:"alipay,omitempty" database:"foreignKey:id"`
	Amount              int                                `json:"amount,omitempty" database:"foreignKey:id"`
	AuBecsDebit         *types.SourceTypeAuBecsDebit       `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	Bancontact          *types.SourceTypeBancontact        `json:"bancontact,omitempty" database:"foreignKey:id"`
	Card                *types.SourceTypeCard              `json:"card,omitempty" database:"foreignKey:id"`
	CardPresent         *types.SourceTypeCardPresent       `json:"card_present,omitempty" database:"foreignKey:id"`
	ClientSecret        string                             `json:"client_secret,omitempty"`
	CodeVerification    *SourceCodeVerificationFlow        `json:"code_verification,omitempty" database:"foreignKey:id"`
	Currency            string                             `json:"currency,omitempty"`
	Eps                 *types.SourceTypeEps               `json:"eps,omitempty" database:"foreignKey:id"`
	Flow                string                             `json:"flow"`
	Giropay             *types.SourceTypeGiropay           `json:"giropay,omitempty" database:"foreignKey:id"`
	Ideal               *types.SourceTypeIdeal             `json:"ideal,omitempty" database:"foreignKey:id"`
	Klarna              *types.SourceTypeKlarna            `json:"klarna,omitempty" database:"foreignKey:id"`
	Multibanco          *types.SourceTypeMultibanco        `json:"multibanco,omitempty" database:"foreignKey:id"`
	Owner               *SourceOwner                       `json:"owner,omitempty" database:"foreignKey:id"`
	P24                 *types.SourceTypeP24               `json:"p24,omitempty" database:"foreignKey:id"`
	Receiver            *SourceReceiverFlow                `json:"receiver,omitempty" database:"foreignKey:id"`
	Redirect            *SourceRedirectFlow                `json:"redirect,omitempty" database:"foreignKey:id"`
	SepaDebit           *types.SourceTypeSepaDebit         `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	Sofort              *types.SourceTypeSofort            `json:"sofort,omitempty" database:"foreignKey:id"`
	SourceOrder         *SourceOrder                       `json:"source_order,omitempty" database:"foreignKey:id"`
	StatementDescriptor string                             `json:"statement_descriptor,omitempty"`
	Status              string                             `json:"status"`
	ThreeDSecure        *types.SourceTypeThreeDSecure      `json:"three_d_secure,omitempty" database:"foreignKey:id"`
	Type                string                             `json:"type"`
	Usage               string                             `json:"usage,omitempty"`
	Wechat              *types.SourceTypeWechat            `json:"wechat,omitempty" database:"foreignKey:id"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
