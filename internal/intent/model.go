package intent

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/charge"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent/methods"
	"github.com/driver005/gateway/internal/intent/next"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/paymentFlow"
)

// PaymentIntentCardProcessing
type PaymentIntentCardProcessing struct {
	CustomerNotification *PaymentIntentProcessingCustomerNotification `json:"customer_notification,omitempty"`
}

// PaymentIntentProcessing
type PaymentIntentProcessing struct {
	Card *PaymentIntentCardProcessing `json:"card,omitempty"`
	// Type of the payment method for which payment is in `processing` state, one of `card`.
	Type string `json:"type"`
}

// PaymentIntentProcessingCustomerNotification
type PaymentIntentProcessingCustomerNotification struct {
	// Whether customer approval has been requested for this payment. For payments greater than INR 15000 or mandate amount, the customer must provide explicit approval of the payment with their bank.
	ApprovalRequested bool `json:"approval_requested,omitempty"`
	// If customer approval is required, they need to provide approval before this time.
	CompletesAt int `json:"completes_at,omitempty"`
}

// PaymentIntentReview ID of the review associated with this PaymentIntent, if any.
type PaymentIntentReview struct {
	// Review *Review
}

type PaymentIntent struct {
	core.Model

	Amount                    int                                                           `json:"amount"`
	AmountCapturable          *int                                                          `json:"amount_capturable,omitempty"`
	AmountDetails             *paymentFlow.PaymentFlowsAmountDetails                        `json:"amount_details,omitempty"`
	AmountReceived            *int                                                          `json:"amount_received,omitempty"`
	ApplicationFeeAmount      int                                                           `json:"application_fee_amount,omitempty"`
	AutomaticPaymentMethods   *paymentFlow.PaymentFlowsAutomaticPaymentMethodsPaymentIntent `json:"automatic_payment_methods,omitempty"`
	CanceledAt                int                                                           `json:"canceled_at,omitempty"`
	CancellationReason        string                                                        `json:"cancellation_reason,omitempty"`
	CaptureMethod             string                                                        `json:"capture_method"`
	ClientSecret              string                                                        `json:"client_secret,omitempty"`
	ConfirmationMethod        string                                                        `json:"confirmation_method"`
	Currency                  string                                                        `json:"currency"`
	Customer                  customer.Customer                                             `json:"customer,omitempty"`
	Description               string                                                        `json:"description,omitempty"`
	Invoice                   invoice.Invoice                                               `json:"invoice,omitempty"`
	LastPaymentError          models.ApiErrors                                              `json:"last_payment_error,omitempty"`
	LatestCharge              charge.Charge                                                 `json:"latest_charge,omitempty"`
	NextAction                next.PaymentIntentNextAction1                                 `json:"next_action,omitempty"`
	PaymentMethod             method.PaymentMethod                                          `json:"payment_method,omitempty"`
	PaymentMethodOptions      methods.PaymentIntentPaymentMethodOptions                     `json:"payment_method_options,omitempty"`
	PaymentMethodTypes        []string                                                      `json:"payment_method_types"`
	Processing                PaymentIntentProcessing                                       `json:"processing,omitempty"`
	ReceiptEmail              string                                                        `json:"receipt_email,omitempty"`
	Review                    PaymentIntentReview                                           `json:"review,omitempty"`
	SetupFutureUsage          string                                                        `json:"setup_future_usage,omitempty"`
	Shipping                  address.Shipping                                              `json:"shipping,omitempty"`
	StatementDescriptor       string                                                        `json:"statement_descriptor,omitempty"`
	StatementDescriptorSuffix string                                                        `json:"statement_descriptor_suffix,omitempty"`
	Status                    string                                                        `json:"status"`
	TransferGroup             string                                                        `json:"transfer_group,omitempty"`
}
