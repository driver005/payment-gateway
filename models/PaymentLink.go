package models

type PaymentLink struct {
	ID              string `json:"id"`
	Object          string `json:"object"`
	Active          bool   `json:"active"`
	AfterCompletion struct {
		HostedConfirmation struct {
			CustomMessage interface{} `json:"custom_message"`
		} `json:"hosted_confirmation"`
		Type string `json:"type"`
	} `json:"after_completion"`
	AllowPromotionCodes   bool        `json:"allow_promotion_codes"`
	ApplicationFeeAmount  interface{} `json:"application_fee_amount"`
	ApplicationFeePercent interface{} `json:"application_fee_percent"`
	AutomaticTax          struct {
		Enabled bool `json:"enabled"`
	} `json:"automatic_tax"`
	BillingAddressCollection string `json:"billing_address_collection"`
	Livemode                 bool   `json:"livemode"`
	Metadata                 struct {
	} `json:"metadata"`
	OnBehalfOf            interface{} `json:"on_behalf_of"`
	PaymentMethodTypes    interface{} `json:"payment_method_types"`
	PhoneNumberCollection struct {
		Enabled bool `json:"enabled"`
	} `json:"phone_number_collection"`
	ShippingAddressCollection interface{} `json:"shipping_address_collection"`
	SubscriptionData          interface{} `json:"subscription_data"`
	TransferData              interface{} `json:"transfer_data"`
	URL                       string      `json:"url"`
}
