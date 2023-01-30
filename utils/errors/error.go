package errors

import (
	"github.com/driver005/gateway/core"
)

// ApiErrorsSource The source object for errors returned on a request involving a source.
type ApiErrorsSource struct {
	core.Model

	// BankAccount *bank.BankAccount
	// Card        *card.Card
	// Source      *source.Source
}

// ApiErrors
type ApiErrors struct {
	core.Model

	// For card errors, the ID of the failed charge.
	Charge string `json:"charge,omitempty"`
	// For some errors that could be handled programmatically, a short string indicating the [error code](https://stripe.com/docs/error-codes) reported.
	Code string `json:"code,omitempty"`
	// For card errors resulting from a card issuer decline, a short string indicating the [card issuer's reason for the decline](https://stripe.com/docs/declines#issuer-declines) if they provide one.
	DeclineCode string `json:"decline_code,omitempty"`
	// A URL to more information about the [error code](https://stripe.com/docs/error-codes) reported.
	DocUrl string `json:"doc_url,omitempty"`
	// A human-readable message providing more details about the error. For card errors, these messages can be shown to your users.
	Message string `json:"message,omitempty"`
	// If the error is parameter-specific, the parameter related to the error. For example, you can use this to display a message near the correct form field.
	Param string `json:"param,omitempty"`
	// If the error is specific to the type of payment method, the payment method type that had a problem. This field is only populated for invoice-related errors.
	PaymentMethodType string `json:"payment_method_type,omitempty"`
	// A URL to the request log entry in your dashboard.
	RequestLogUrl string `json:"request_log_url,omitempty"`
	// The type of error returned. One of `api_error`, `card_error`, `idempotency_error`, or `invalid_request_error`
	Type string `json:"type,omitempty"`
}
