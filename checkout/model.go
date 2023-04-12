package checkout

import (
	"github.com/driver005/gateway/billing/invoice"
	"github.com/driver005/gateway/billing/subscription"
	"github.com/driver005/gateway/checkout/options"
	"github.com/driver005/gateway/checkout/page"
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/intent"
	setup_intent "github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/link"
	"github.com/driver005/gateway/products/item"
	"github.com/driver005/gateway/utils/address"
	"github.com/lib/pq"
)

type CheckoutSession struct {
	core.Model

	AfterExpiration           *page.PaymentPagesCheckoutSessionAfterExpiration           `json:"after_expiration,omitempty" database:"foreignKey:id"`
	AllowPromotionCodes       bool                                                       `json:"allow_promotion_codes,omitempty"`
	AmountSubtotal            int                                                        `json:"amount_subtotal,omitempty"`
	AmountTotal               int                                                        `json:"amount_total,omitempty"`
	AutomaticTax              bool                                                       `json:"automatic_tax,omitempty"`
	BillingAddressCollection  BillingAddressCollection                                   `json:"billing_address_collection,omitempty"`
	CancelUrl                 string                                                     `json:"cancel_url,omitempty"`
	ClientReferenceId         string                                                     `json:"client_reference_id,omitempty"`
	Consent                   *page.PaymentPagesCheckoutSessionConsent                   `json:"consent,omitempty" database:"foreignKey:id"`
	ConsentCollection         *page.PaymentPagesCheckoutSessionConsentCollection         `json:"consent_collection,omitempty" database:"foreignKey:id"`
	Currency                  string                                                     `json:"currency,omitempty"`
	CustomText                *page.PaymentPagesCheckoutSessionCustomText                `json:"custom_text,omitempty" database:"foreignKey:id"`
	Customer                  *customer.Customer                                         `json:"customer,omitempty" database:"foreignKey:id"`
	CustomerCreation          CustomerCreation                                           `json:"customer_creation,omitempty"`
	CustomerDetails           *page.PaymentPagesCheckoutSessionCustomerDetails           `json:"customer_details,omitempty" database:"foreignKey:id"`
	CustomerEmail             string                                                     `json:"customer_email,omitempty"`
	ExpiresAt                 int                                                        `json:"expires_at,omitempty"`
	Invoice                   *invoice.Invoice                                           `json:"invoice,omitempty" database:"foreignKey:id"`
	InvoiceCreation           *page.PaymentPagesCheckoutSessionInvoiceCreation           `json:"invoice_creation,omitempty" database:"foreignKey:id"`
	LineItems                 []item.LineItem                                            `json:"line_items,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	Locale                    Locale                                                     `json:"locale,omitempty"`
	Mode                      Mode                                                       `json:"mode,omitempty"`
	PaymentIntent             *intent.PaymentIntent                                      `json:"payment_intent,omitempty" database:"foreignKey:id"`
	PaymentLink               *link.PaymentLink                                          `json:"payment_link,omitempty" database:"foreignKey:id"`
	PaymentMethodCollection   PaymentMethodCollection                                    `json:"payment_method_collection,omitempty"`
	PaymentMethodOptions      *options.CheckoutSessionPaymentMethodOptions               `json:"payment_method_options,omitempty" database:"foreignKey:id"`
	PaymentMethodTypes        pq.StringArray                                             `json:"payment_method_types,omitempty" database:"type:varchar(64)[]"`
	PaymentStatus             PaymentStatus                                              `json:"payment_status,omitempty"`
	PhoneNumberCollection     bool                                                       `json:"phone_number_collection,omitempty"`
	RecoveredFrom             string                                                     `json:"recovered_from,omitempty"`
	SetupIntent               *setup_intent.SetupIntent                                  `json:"setup_intent,omitempty" database:"foreignKey:id"`
	ShippingAddressCollection *page.PaymentPagesCheckoutSessionShippingAddressCollection `json:"shipping_address_collection,omitempty" database:"foreignKey:id"`
	ShippingCost              *page.PaymentPagesCheckoutSessionShippingCost              `json:"shipping_cost,omitempty" database:"foreignKey:id"`
	ShippingDetails           *address.Shipping                                          `json:"shipping_details,omitempty" database:"foreignKey:id"`
	ShippingOptions           []page.PaymentPagesCheckoutSessionShippingOption           `json:"shipping_options,omitempty" database:"foreignKey:id" swaggertype:"array,string" format:"uuid"`
	Status                    Status                                                     `json:"status,omitempty"`
	SubmitType                SubmitType                                                 `json:"submit_type,omitempty"`
	Subscription              *subscription.Subscription                                 `json:"subscription,omitempty" database:"foreignKey:id"`
	SuccessUrl                string                                                     `json:"success_url,omitempty"`
	TaxIdCollection           bool                                                       `json:"tax_id_collection,omitempty"`
	TotalDetails              *page.PaymentPagesCheckoutSessionTotalDetails              `json:"total_details,omitempty" database:"foreignKey:id"`
	Url                       string                                                     `json:"url,omitempty"`
}
