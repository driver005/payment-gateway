package link

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/link/options"
	"github.com/driver005/gateway/products/item"
	"github.com/lib/pq"
)

type PaymentLink struct {
	core.Model

	Active                    bool                                                   `json:"active,omitempty"`
	AfterCompletion           *options.PaymentLinksResourceAfterCompletion           `json:"after_completion,omitempty" database:"foreignKey:id"`
	AllowPromotionCodes       bool                                                   `json:"allow_promotion_codes,omitempty"`
	ApplicationFeeAmount      int                                                    `json:"application_fee_amount,omitempty"`
	ApplicationFeePercent     float64                                                `json:"application_fee_percent,omitempty"`
	AutomaticTax              bool                                                   `json:"automatic_tax,omitempty"`
	BillingAddressCollection  string                                                 `json:"billing_address_collection,omitempty"`
	ConsentCollection         *options.PaymentLinksResourceConsentCollection         `json:"consent_collection,omitempty" database:"foreignKey:id"`
	Currency                  string                                                 `json:"currency,omitempty"`
	CustomText                *options.PaymentLinksResourceCustomText                `json:"custom_text,omitempty" database:"foreignKey:id"`
	CustomerCreation          string                                                 `json:"customer_creation,omitempty"`
	LineItems                 []item.LineItem                                        `json:"line_items,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	PaymentIntentData         *options.PaymentLinksResourcePaymentIntentData         `json:"payment_intent_data,omitempty" database:"foreignKey:id"`
	PaymentMethodCollection   string                                                 `json:"payment_method_collection,omitempty"`
	PaymentMethodTypes        pq.StringArray                                         `json:"payment_method_types,omitempty" database:"type:varchar(64)[]"`
	PhoneNumberCollection     bool                                                   `json:"phone_number_collection,omitempty"`
	ShippingAddressCollection *options.PaymentLinksResourceShippingAddressCollection `json:"shipping_address_collection,omitempty" database:"foreignKey:id"`
	ShippingOptions           []options.PaymentLinksResourceShippingOption           `json:"shipping_options,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	SubmitType                string                                                 `json:"submit_type,omitempty"`
	SubscriptionData          *options.PaymentLinksResourceSubscriptionData          `json:"subscription_data,omitempty" database:"foreignKey:id"`
	TaxIdCollection           bool                                                   `json:"tax_id_collection,omitempty"`
	Url                       string                                                 `json:"url,omitempty"`
}
