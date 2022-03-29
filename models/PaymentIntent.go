package models

type PaymentIntent struct {
	ID                      string      `json:"id"`
	Object                  string      `json:"object"`
	Amount                  int         `json:"amount"`
	AmountCapturable        int         `json:"amount_capturable"`
	AmountReceived          int         `json:"amount_received"`
	Application             interface{} `json:"application"`
	ApplicationFeeAmount    interface{} `json:"application_fee_amount"`
	AutomaticPaymentMethods interface{} `json:"automatic_payment_methods"`
	CanceledAt              interface{} `json:"canceled_at"`
	CancellationReason      interface{} `json:"cancellation_reason"`
	CaptureMethod           string      `json:"capture_method"`
	Charges                 struct {
		Object  string        `json:"object"`
		Data    []interface{} `json:"data"`
		HasMore bool          `json:"has_more"`
		URL     string        `json:"url"`
	} `json:"charges"`
	ClientSecret       string      `json:"client_secret"`
	ConfirmationMethod string      `json:"confirmation_method"`
	Created            int         `json:"created"`
	Currency           string      `json:"currency"`
	Customer           interface{} `json:"customer"`
	Description        interface{} `json:"description"`
	Invoice            interface{} `json:"invoice"`
	LastPaymentError   interface{} `json:"last_payment_error"`
	Livemode           bool        `json:"livemode"`
	Metadata           struct {
	} `json:"metadata"`
	NextAction           interface{} `json:"next_action"`
	OnBehalfOf           interface{} `json:"on_behalf_of"`
	PaymentMethod        interface{} `json:"payment_method"`
	PaymentMethodOptions struct {
	} `json:"payment_method_options"`
	PaymentMethodTypes        []string    `json:"payment_method_types"`
	Processing                interface{} `json:"processing"`
	ReceiptEmail              interface{} `json:"receipt_email"`
	Redaction                 interface{} `json:"redaction"`
	Review                    interface{} `json:"review"`
	SetupFutureUsage          interface{} `json:"setup_future_usage"`
	Shipping                  interface{} `json:"shipping"`
	StatementDescriptor       interface{} `json:"statement_descriptor"`
	StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	Status                    string      `json:"status"`
	TransferData              interface{} `json:"transfer_data"`
	TransferGroup             interface{} `json:"transfer_group"`
}