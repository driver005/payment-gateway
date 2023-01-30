package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

type PaymentIntent struct {
	core.Model

	Amount                  int       `json:"amount"`
	AmountCapturable        int       `json:"amount_capturable"`
	AmountReceived          int       `json:"amount_received"`
	Tip                     int       `json:"tip"`
	AutomaticPaymentMethods bool      `json:"automatic_payment_methods"`
	CanceledAt              time.Time `json:"canceled_at"`
	CancellationReason      string    `json:"cancellation_reason"`
	CaptureMethod           string    `json:"capture_method"`
	// Charges                 struct {
	// 	Object  string        `json:"object"`
	// 	Data    []interface{} `json:"data"`
	// 	HasMore bool          `json:"has_more"`
	// 	URL     string        `json:"url"`
	// } `json:"charges"`
	ClientSecret       string    `json:"client_secret"`
	ConfirmationMethod string    `json:"confirmation_method"`
	Currency           *Currency `json:"currency" database:"foreignKey:id"`
	Customer           *Customer `json:"customer" database:"foreignKey:id"`
	Description        string    `json:"description"`
	// Invoice            interface{} `json:"invoice"`
	// LastPaymentError   interface{} `json:"last_payment_error"`
	// NextAction           interface{} `json:"next_action"`
	PaymentMethod *PaymentMethod `json:"payment_method" database:"foreignKey:id"`
	// PaymentMethodOptions struct{}    `json:"payment_method_options"`
	PaymentMethodTypes []string    `json:"payment_method_types" database:"type:text[]"`
	Processing         interface{} `json:"processing"`
	ReceiptEmail       string      `json:"receipt_email"`
	// Redaction                 interface{} `json:"redaction"`
	// Review                    interface{} `json:"review"`
	// SetupFutureUsage          interface{} `json:"setup_future_usage"`
	// Shipping                  interface{} `json:"shipping"`
	// StatementDescriptor       interface{} `json:"statement_descriptor"`
	// StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	Status string `json:"status"`
}
