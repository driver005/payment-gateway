package intent

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent/methods"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/errors"
	"github.com/driver005/gateway/utils/paymentFlow"
	"github.com/driver005/gateway/utils/review"
	"github.com/google/uuid"
)

// PaymentIntentCardProcessing
type PaymentIntentCardProcessing struct {
	core.Model

	CustomerNotification *PaymentIntentProcessingCustomerNotification `json:"customer_notification,omitempty" database:"foreignKey:id"`
}

// PaymentIntentProcessing
type PaymentIntentProcessing struct {
	core.Model

	Card *PaymentIntentCardProcessing `json:"card,omitempty" database:"foreignKey:id"`
	// Type of the payment method for which payment is in `processing` state, one of `card`.
	Type string `json:"type,omitempty"`
}

// PaymentIntentProcessingCustomerNotification
type PaymentIntentProcessingCustomerNotification struct {
	core.Model

	// Whether customer approval has been requested for this payment. For payments greater than INR 15000 or mandate amount, the customer must provide explicit approval of the payment with their bank.
	ApprovalRequested bool `json:"approval_requested,omitempty"`
	// If customer approval is required, they need to provide approval before this time.
	CompletesAt int `json:"completes_at,omitempty"`
}

type PaymentIntent struct {
	core.Model

	Amount                  int                                                           `json:"amount,omitempty"`
	AmountCapturable        int                                                           `json:"amount_capturable,omitempty"`
	AmountDetails           *paymentFlow.PaymentFlowsAmountDetails                        `json:"amount_details,omitempty" database:"foreignKey:id"`
	AmountReceived          int                                                           `json:"amount_received,omitempty"`
	AutomaticPaymentMethods *paymentFlow.PaymentFlowsAutomaticPaymentMethodsPaymentIntent `json:"automatic_payment_methods,omitempty" database:"foreignKey:id"`
	CanceledAt              int                                                           `json:"canceled_at,omitempty"`
	CancellationReason      string                                                        `json:"cancellation_reason,omitempty"`
	CaptureMethod           string                                                        `json:"capture_method,omitempty"`
	ClientSecret            string                                                        `json:"client_secret,omitempty"`
	ConfirmationMethod      string                                                        `json:"confirmation_method,omitempty"`
	Currency                string                                                        `json:"currency,omitempty"`

	CustomerId uuid.UUID          `json:"customer_id" database:"default:null"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`

	Description string `json:"description,omitempty"`
	// Invoice          *invoice.Invoice              `json:"invoice,omitempty" database:"foreignKey:id"`
	LastPaymentError *errors.ApiErrors             `json:"last_payment_error,omitempty" database:"foreignKey:id"`
	LatestCharge     *charge.Charge                `json:"latest_charge,omitempty" database:"foreignKey:id"`
	NextAction       *next.PaymentIntentNextAction `json:"next_action,omitempty" database:"foreignKey:id"`

	PaymentMethodId uuid.UUID             `json:"payment_method_id" database:"default:null"`
	PaymentMethod   *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id"`

	PaymentMethodOptions      *methods.PaymentIntentPaymentMethodOptions `json:"payment_method_options,omitempty" database:"foreignKey:id"`
	PaymentMethodTypes        []string                                   `json:"payment_method_types,omitempty" database:"type:text[]"`
	Processing                *PaymentIntentProcessing                   `json:"processing,omitempty" database:"foreignKey:id"`
	ReceiptEmail              string                                     `json:"receipt_email,omitempty"`
	Review                    *review.Review                             `json:"review,omitempty" database:"foreignKey:id"`
	SetupFutureUsage          string                                     `json:"setup_future_usage,omitempty"`
	Shipping                  *address.Shipping                          `json:"shipping,omitempty" database:"foreignKey:id"`
	StatementDescriptor       string                                     `json:"statement_descriptor,omitempty"`
	StatementDescriptorSuffix string                                     `json:"statement_descriptor_suffix,omitempty"`
	Status                    string                                     `json:"status,omitempty"`
}
