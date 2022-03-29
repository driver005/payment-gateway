package models

type Session struct {
	ID                  string      `json:"id"`
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
	Metadata                 struct {
	} `json:"metadata"`
	Mode                 string      `json:"mode"`
	PaymentIntent        string      `json:"payment_intent"`
	PaymentLink          interface{} `json:"payment_link"`
	PaymentMethodOptions struct {
	} `json:"payment_method_options"`
	PaymentMethodTypes    []string `json:"payment_method_types"`
	PaymentStatus         string   `json:"payment_status"`
	PhoneNumberCollection struct {
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