package models

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/gofrs/uuid"
)

type Refund struct {
	ID                        uuid.UUID `json:"id"`
	Object                    string    `json:"object"`
	Amount                    int       `json:"amount"`
	BalanceTransaction        string    `json:"balance_transaction"`
	Charge                    uuid.UUID `json:"charge"`
	Currency                  string    `json:"currency"`
	Description               string    `json:"description"`
	FailureBalanceTransaction string    `json:"failure_balance_transaction"`
	FailureReason             string    `json:"failure_reason"`
	Metadata                  uuid.UUID `json:"metadata"`
	NextAction                struct {
		DisplayDetails struct {
			EmailSent struct {
				SentAt string `json:"sent_at"`
				SentTo string `json:"sent_to"`
			} `json:"email_sent"`
			ExpiresAt string `json:"expires_at"`
		} `json:"display_details"`
		Type string `json:"type"`
	} `json:"next_action"`
	PaymentIntent          uuid.UUID `json:"payment_intent"`
	Reason                 string    `json:"reason"`
	ReceiptNumber          string    `json:"receipt_number"`
	SourceTransferReversal string    `json:"source_transfer_reversal"`
	Status                 string    `json:"status"`
	TransferReversal       string    `json:"transfer_reversal"`
}

func (r Refund) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.ID, validation.Required, is.UUID),
		validation.Field(&r.Object, validation.Required, validation.In("Refund")),
		validation.Field(&r.Amount, validation.Required),
		validation.Field(&r.BalanceTransaction),
		validation.Field(&r.Charge, validation.Required, is.UUID),
		validation.Field(&r.Currency, validation.Required, is.CurrencyCode),
		validation.Field(&r.Description),
		validation.Field(&r.FailureBalanceTransaction),
		validation.Field(&r.FailureReason, validation.In("lost_or_stolen_card", "expired_or_canceled_card", "unknown")),
		validation.Field(&r.Metadata, is.UUID),
		validation.Field(&r.NextAction.DisplayDetails.EmailSent.SentTo, is.Email),
		validation.Field(&r.NextAction.DisplayDetails.EmailSent.SentAt, validation.Date(time.ANSIC)),
		validation.Field(&r.NextAction.DisplayDetails.ExpiresAt, validation.Date(time.ANSIC)),
		validation.Field(&r.NextAction.Type),
		validation.Field(&r.PaymentIntent, is.UUID),
		validation.Field(&r.Reason, validation.In("duplicate", "fraudulent", "requested_by_customer", "expired_uncaptured_charge")),
		validation.Field(&r.ReceiptNumber),
		validation.Field(&r.SourceTransferReversal),
		validation.Field(&r.Status, validation.In("pending", "succeeded", "failed", "canceled")),
	)
}
