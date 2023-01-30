package charge

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/method/details"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/radar"
	"github.com/driver005/gateway/utils/review"
	"github.com/driver005/gateway/utils/rule"
	"github.com/google/uuid"
)

// ChargeFraudDetails
type ChargeFraudDetails struct {
	core.Model

	// Assessments from Stripe. If set, the value is `fraudulent`.
	StripeReport string `json:"stripe_report,omitempty"`
	// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
	UserReport string `json:"user_report,omitempty"`
}

// ChargeOutcome
type ChargeOutcome struct {
	core.Model

	// Possible values are `approved_by_network`, `declined_by_network`, `not_sent_to_network`, and `reversed_after_approval`. The value `reversed_after_approval` indicates the payment was [blocked by Stripe](https://stripe.com/docs/declines#blocked-payments) after bank authorization, and may temporarily appear as \"pending\" on a cardholder's statement.
	NetworkStatus string `json:"network_status,omitempty"`
	// An enumerated value providing a more detailed explanation of the outcome's `type`. Charges blocked by Radar's default block rule have the value `highest_risk_level`. Charges placed in review by Radar's default review rule have the value `elevated_risk_level`. Charges authorized, blocked, or placed in review by custom rules have the value `rule`. See [understanding declines](https://stripe.com/docs/declines) for more details.
	Reason string `json:"reason,omitempty"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are `normal`, `elevated`, `highest`. For non-card payments, and card-based payments predating the public assignment of risk levels, this field will have the value `not_assessed`. In the event of an error in the evaluation, this field will have the value `unknown`. This field is only available with Radar.
	RiskLevel string `json:"risk_level,omitempty"`
	// Stripe Radar's evaluation of the riskiness of the payment. Possible values for evaluated payments are between 0 and 100. For non-card payments, card-based payments predating the public assignment of risk scores, or in the event of an error during evaluation, this field will not be present. This field is only available with Radar for Fraud Teams.
	RiskScore int        `json:"risk_score,omitempty"`
	Rule      *rule.Rule `json:"rule,omitempty" database:"foreignKey:id"`
	// A human-readable description of the outcome type and reason, designed for you (the recipient of the payment), not your customer.
	SellerMessage string `json:"seller_message,omitempty"`
	// Possible values are `authorized`, `manual_review`, `issuer_declined`, `blocked`, and `invalid`. See [understanding declines](https://stripe.com/docs/declines) and [Radar reviews](https://stripe.com/docs/radar/reviews) for details.
	Type string `json:"type,omitempty"`
}

type Charge struct {
	core.Model

	Amount                        int                           `json:"amount,omitempty"`
	AmountCaptured                int                           `json:"amount_captured,omitempty"`
	AmountRefunded                int                           `json:"amount_refunded,omitempty"`
	BillingDetails                *address.BillingDetails       `json:"billing_details,omitempty" database:"foreignKey:id"`
	CalculatedStatementDescriptor string                        `json:"calculated_statement_descriptor,omitempty"`
	Captured                      bool                          `json:"captured,omitempty"`
	Currency                      string                        `json:"currency,omitempty"`
	Description                   string                        `json:"description,omitempty"`
	Disputed                      bool                          `json:"disputed,omitempty"`
	FailureCode                   string                        `json:"failure_code,omitempty"`
	FailureMessage                string                        `json:"failure_message,omitempty"`
	FraudDetails                  *ChargeFraudDetails           `json:"fraud_details,omitempty" database:"foreignKey:id"`
	Outcome                       *ChargeOutcome                `json:"outcome,omitempty"  database:"foreignKey:id"`
	Paid                          bool                          `json:"paid"`
	PaymentMethodDetails          *details.PaymentMethodDetails `json:"payment_method_details,omitempty" database:"foreignKey:id"`
	RadarOptions                  *radar.RadarOptions           `json:"radar_options,omitempty" database:"foreignKey:id"`
	ReceiptEmail                  string                        `json:"receipt_email,omitempty"`
	ReceiptNumber                 string                        `json:"receipt_number,omitempty"`
	ReceiptUrl                    string                        `json:"receipt_url,omitempty"`
	Refunded                      bool                          `json:"refunded,omitempty"`
	Shipping                      *address.Shipping             `json:"shipping,omitempty" database:"foreignKey:id"`
	StatementDescriptor           string                        `json:"statement_descriptor,omitempty"`
	StatementDescriptorSuffix     string                        `json:"statement_descriptor_suffix,omitempty"`
	Status                        string                        `json:"status,omitempty"`

	BalanceTransactionId        uuid.UUID                   `json:"balance_transaction_id,omitempty"`
	BalanceTransaction          *balance.BalanceTransaction `json:"balance_transaction,omitempty" database:"foreignKey:balance_transaction_id"`
	CustomerId                  uuid.UUID                   `json:"customer_id" database:"default:null"`
	Customer                    *customer.Customer          `json:"customer,omitempty" database:"foreignKey:customer_id"`
	FailureBalanceTransactionId uuid.UUID                   `json:"failure_balance_transaction_id,omitempty"`
	FailureBalanceTransaction   *balance.BalanceTransaction `json:"failure_balance_transaction,omitempty" database:"foreignKey:failure_balance_transaction_id"`
	PaymentMethodId             uuid.UUID                   `json:"payment_method_id" database:"default:null"`
	PaymentMethod               *method.PaymentMethod       `json:"payment_method,omitempty" database:"foreignKey:payment_method_id"`
	ReviewId                    uuid.UUID                   `json:"review_id,omitempty"`
	Review                      *review.Review              `json:"review,omitempty" database:"foreignKey:review_id"`
	// PaymentIntentId uuid.UUID             `json:"payment_intent_id,omitempty"`
	// PaymentIntent   *intent.PaymentIntent `json:"payment_intent,omitempty" database:"foreignKey:payment_intent_id"`
}
