package models

import (
	"github.com/driver005/database"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

type PaymentIntent struct {
	Id                 uuid.UUID   `json:"id"`
	Amount             int         `json:"amount"`
	AmountReceived     int         `json:"amount_received"`
	CanceledAt         interface{} `json:"canceled_at"`
	CancellationReason interface{} `json:"cancellation_reason"`
	Charges            struct {
		Object  string        `json:"object"`
		Data    []interface{} `json:"data"`
		HasMore bool          `json:"has_more"`
		URL     string        `json:"url"`
	} `json:"charges"`
	ConfirmationMethod   string      `json:"confirmation_method"`
	Currency             string      `json:"currency"`
	Description          interface{} `json:"description"`
	Invoice              interface{} `json:"invoice"`
	NextAction           interface{} `json:"next_action"`
	PaymentMethod        interface{} `json:"payment_method"`
	PaymentMethodOptions struct{}    `json:"payment_method_options"`
	ReceiptEmail         interface{} `json:"receipt_email"`
	StatementDescriptor  interface{} `json:"statement_descriptor"`
	Status               string      `json:"status"`
	Created              int         `json:"created"`
	// Application             interface{} `json:"application"`
	// ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
	// AutomaticPaymentMethods interface{} `json:"automatic_payment_methods"`
	// CaptureMethod           string      `json:"capture_method"`
	// ClientSecret       string      `json:"client_secret"`
	// Customer           interface{} `json:"customer"`
	// LastPaymentError   interface{} `json:"last_payment_error"`
	// Livemode           bool        `json:"livemode"`
	// Metadata           struct {} `json:"metadata"`
	// OnBehalfOf           interface{} `json:"on_behalf_of"`
	// PaymentMethodTypes        []string    `json:"payment_method_types"`
	// Processing                interface{} `json:"processing"`
	// Redaction                 interface{} `json:"redaction"`
	// Review                    interface{} `json:"review"`
	// SetupFutureUsage          interface{} `json:"setup_future_usage"`
	// Shipping                  interface{} `json:"shipping"`
	// StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	// TransferData              interface{} `json:"transfer_data"`
	// TransferGroup             interface{} `json:"transfer_group"`
}

func (p *PaymentIntent) BeforeCreate(tx *database.DB) (err error) {
	p.Id, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

func (p PaymentIntent) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Id, validation.Required, is.UUID),
		validation.Field(&p.Amount, validation.Required, validation.Max(99999999)),
		// validation.Field(&p.Address1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		// validation.Field(&p.ZipCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
	)
}
