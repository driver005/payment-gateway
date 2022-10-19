package models

import (
	"github.com/driver005/database"
	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/gofrs/uuid"
)

type Session struct {
	ID                  uuid.UUID   `json:"id"`
	Object              string      `json:"object"`
	AfterExpiration     interface{} `json:"after_expiration"`
	AllowPromotionCodes interface{} `json:"allow_promotion_codes"`
	AmountSubtotal      interface{} `json:"amount_subtotal"`
	AmountTotal         interface{} `json:"amount_total"`
	AutomaticTax        struct {
		Enabled bool        `json:"enabled"`
		Status  interface{} `json:"status"`
	} `json:"automatic_tax"`
	BillingAddressCollection interface{} `json:"billing_address_collection"`
	CancelURL                string      `json:"cancel_url"`
	ClientReferenceID        interface{} `json:"client_reference_id"`
	Consent                  interface{} `json:"consent"`
	ConsentCollection        interface{} `json:"consent_collection"`
	Currency                 interface{} `json:"currency"`
	Customer                 interface{} `json:"customer"`
	CustomerCreation         interface{} `json:"customer_creation"`
	CustomerDetails          interface{} `json:"customer_details"`
	CustomerEmail            interface{} `json:"customer_email"`
	ExpiresAt                int         `json:"expires_at"`
	Livemode                 bool        `json:"livemode"`
	Locale                   interface{} `json:"locale"`
	Metadata                 struct{}    `json:"metadata"`
	Mode                     string      `json:"mode"`
	PaymentIntent            string      `json:"payment_intent"`
	PaymentLink              interface{} `json:"payment_link"`
	PaymentMethodOptions     struct{}    `json:"payment_method_options"`
	PaymentMethodTypes       []string    `json:"payment_method_types"`
	PaymentStatus            string      `json:"payment_status"`
	PhoneNumberCollection    struct {
		Enabled bool `json:"enabled"`
	} `json:"phone_number_collection"`
	RecoveredFrom             interface{}   `json:"recovered_from"`
	SetupIntent               interface{}   `json:"setup_intent"`
	Shipping                  interface{}   `json:"shipping"`
	ShippingAddressCollection interface{}   `json:"shipping_address_collection"`
	ShippingOptions           []interface{} `json:"shipping_options"`
	ShippingRate              interface{}   `json:"shipping_rate"`
	Status                    string        `json:"status"`
	SubmitType                interface{}   `json:"submit_type"`
	Subscription              interface{}   `json:"subscription"`
	SuccessURL                string        `json:"success_url"`
	TotalDetails              interface{}   `json:"total_details"`
	URL                       interface{}   `json:"url"`
}

func (s *Session) BeforeCreate(tx *database.DB) (err error) {
	s.ID, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}

func (s Session) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.ID, validation.Required, is.UUID),
		validation.Field(&s.Object, validation.Required, validation.In("checkout.session")),

		// validation.Field(&s.Address, validation.When(s.Address.Validate() != nil)),
		// validation.Field(&s.Brand, validation.Required, validation.In("American Express", "Diners Club", "Discover", "JCB", "MasterCard", "UnionPay", "Visa", "Unknown")),
		// validation.Field(&s.Country, validation.Required, is.CountryCode2),
		// validation.Field(&s.Customer, is.UUID),
		// validation.Field(&s.CvcCheck, validation.Required, validation.In("pass", "fail", "unavailable", "unchecked")),
		// validation.Field(&s.DynamicLast4),
		// validation.Field(&s.ExpMonth, validation.Required, validation.Min(1), validation.Max(12)),
		// validation.Field(&s.ExpYear, validation.Required, validation.Min(2000)),
		// validation.Field(&s.Fingerprint, validation.Required, validation.Length(16, 32)),
		// validation.Field(&s.Funding, validation.Required, validation.In("credit", "debit", "prepaid", "unknown")),
		// validation.Field(&s.Last4, validation.Required, validation.Length(4, 4)),
		// validation.Field(&s.Name),
		// validation.Field(&s.TokenizationMethod, validation.Required, validation.In("android_pay", "apple_pay", "masterpass", "visa_checkout")),
	)
}
