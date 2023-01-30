package dispute

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/balance"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/file"
	"github.com/driver005/gateway/internal/intent"
	"github.com/google/uuid"
)

// Dispute A dispute occurs when a customer questions your charge with their card issuer. When this happens, you're given the opportunity to respond to the dispute with evidence that shows that the charge is legitimate. You can find more information about the dispute process in our [Disputes and Fraud](/docs/disputes) documentation.  Related guide: [Disputes and Fraud](https://stripe.com/docs/disputes).
type Dispute struct {
	core.Model

	// Disputed amount. Usually the amount of the charge, but can differ (usually because of currency fluctuation or because only part of the order is disputed).
	Amount int `json:"amount,omitempty"`
	// List of zero, one, or two balance transactions that show funds withdrawn and reinstated to your Stripe account as a result of this dispute.
	BalanceTransactions []balance.BalanceTransaction `json:"balance_transactions,omitempty" database:"foreignKey:id"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency        string                 `json:"currency,omitempty"`
	Evidence        DisputeEvidence        `json:"evidence,omitempty" database:"foreignKey:id"`
	EvidenceDetails DisputeEvidenceDetails `json:"evidence_details,omitempty" database:"foreignKey:id"`
	// If true, it is still possible to refund the disputed payment. Once the payment has been fully refunded, no further funds will be withdrawn from your Stripe account as a result of this dispute.
	IsChargeRefundable bool `json:"is_charge_refundable,omitempty"`
	// Reason given by cardholder for dispute. Possible values are `bank_cannot_process`, `check_returned`, `credit_not_processed`, `customer_initiated`, `debit_not_authorized`, `duplicate`, `fraudulent`, `general`, `incorrect_account_details`, `insufficient_funds`, `product_not_received`, `product_unacceptable`, `subscription_canceled`, or `unrecognized`. Read more about [dispute reasons](https://stripe.com/docs/disputes/categories).
	Reason string `json:"reason,omitempty"`
	// Current status of dispute. Possible values are `warning_needs_response`, `warning_under_review`, `warning_closed`, `needs_response`, `under_review`, `charge_refunded`, `won`, or `lost`.
	Status string `json:"status,omitempty"`

	ChargeId        uuid.UUID             `json:"charge_id,omitempty"`
	Charge          *charge.Charge        `json:"charge,omitempty" database:"foreignKey:charge_id"`
	PaymentIntentId uuid.UUID             `json:"payment_intent_id,omitempty"`
	PaymentIntent   *intent.PaymentIntent `json:"payment_intent,omitempty" database:"foreignKey:payment_intent_id"`
}

// DisputeEvidenceDetails
type DisputeEvidenceDetails struct {
	core.Model

	// Date by which evidence must be submitted in order to successfully challenge dispute. Will be null if the customer's bank or credit card company doesn't allow a response for this particular dispute.
	DueBy int `json:"due_by,omitempty"`
	// Whether evidence has been staged for this dispute.
	HasEvidence bool `json:"has_evidence"`
	// Whether the last evidence submission was submitted past the due date. Defaults to `false` if no evidence submissions have occurred. If `true`, then delivery of the latest evidence is *not* guaranteed.
	PastDue bool `json:"past_due"`
	// The number of times evidence has been submitted. Typically, you may only submit evidence once.
	SubmissionCount int `json:"submission_count"`
}

// DisputeEvidence
type DisputeEvidence struct {
	core.Model

	// Any server or activity logs showing proof that the customer accessed or downloaded the purchased digital product. This information should include IP addresses, corresponding timestamps, and any detailed recorded activity.
	AccessActivityLog string `json:"access_activity_log,omitempty"`
	// The billing address provided by the customer.
	BillingAddress     string     `json:"billing_address,omitempty"`
	CancellationPolicy *file.File `json:"cancellation_policy,omitempty" database:"foreignKey:id"`
	// An explanation of how and when the customer was shown your refund policy prior to purchase.
	CancellationPolicyDisclosure string `json:"cancellation_policy_disclosure,omitempty"`
	// A justification for why the customer's subscription was not canceled.
	CancellationRebuttal  string     `json:"cancellation_rebuttal,omitempty"`
	CustomerCommunication *file.File `json:"customer_communication,omitempty" database:"foreignKey:id"`
	// The email address of the customer.
	CustomerEmailAddress string `json:"customer_email_address,omitempty"`
	// The name of the customer.
	CustomerName string `json:"customer_name,omitempty"`
	// The IP address that the customer used when making the purchase.
	CustomerPurchaseIp           string     `json:"customer_purchase_ip,omitempty"`
	CustomerSignature            *file.File `json:"customer_signature,omitempty" database:"foreignKey:id"`
	DuplicateChargeDocumentation *file.File `json:"duplicate_charge_documentation,omitempty" database:"foreignKey:id"`
	// An explanation of the difference between the disputed charge versus the prior charge that appears to be a duplicate.
	DuplicateChargeExplanation string `json:"duplicate_charge_explanation,omitempty"`
	// The Stripe ID for the prior charge which appears to be a duplicate of the disputed charge.
	DuplicateChargeId string `json:"duplicate_charge_id,omitempty"`
	// A description of the product or service that was sold.
	ProductDescription string     `json:"product_description,omitempty"`
	Receipt            *file.File `json:"receipt,omitempty" database:"foreignKey:id"`
	RefundPolicy       *file.File `json:"refund_policy,omitempty" database:"foreignKey:id"`
	// Documentation demonstrating that the customer was shown your refund policy prior to purchase.
	RefundPolicyDisclosure string `json:"refund_policy_disclosure,omitempty"`
	// A justification for why the customer is not entitled to a refund.
	RefundRefusalExplanation string `json:"refund_refusal_explanation,omitempty"`
	// The date on which the customer received or began receiving the purchased service, in a clear human-readable format.
	ServiceDate          string     `json:"service_date,omitempty"`
	ServiceDocumentation *file.File `json:"service_documentation,omitempty" database:"foreignKey:id"`
	// The address to which a physical product was shipped. You should try to include as complete address information as possible.
	ShippingAddress string `json:"shipping_address,omitempty"`
	// The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc. If multiple carriers were used for this purchase, please separate them with commas.
	ShippingCarrier string `json:"shipping_carrier,omitempty"`
	// The date on which a physical product began its route to the shipping address, in a clear human-readable format.
	ShippingDate          string     `json:"shipping_date,omitempty"`
	ShippingDocumentation *file.File `json:"shipping_documentation,omitempty" database:"foreignKey:id"`
	// The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.
	ShippingTrackingNumber string     `json:"shipping_tracking_number,omitempty"`
	UncategorizedFile      *file.File `json:"uncategorized_file,omitempty" database:"foreignKey:id"`
	// Any additional evidence or statements.
	UncategorizedText string `json:"uncategorized_text,omitempty"`
}
