package types

import "github.com/driver005/gateway/core"

// SourceTypeKlarna struct for SourceTypeKlarna
type SourceTypeKlarna struct {
	core.Model

	BackgroundImageUrl              string `json:"background_image_url,omitempty"`
	ClientToken                     string `json:"client_token,omitempty"`
	FirstName                       string `json:"first_name,omitempty"`
	LastName                        string `json:"last_name,omitempty"`
	Locale                          string `json:"locale,omitempty"`
	LogoUrl                         string `json:"logo_url,omitempty"`
	PageTitle                       string `json:"page_title,omitempty"`
	PayLaterAssetUrlsDescriptive    string `json:"pay_later_asset_urls_descriptive,omitempty"`
	PayLaterAssetUrlsStandard       string `json:"pay_later_asset_urls_standard,omitempty"`
	PayLaterName                    string `json:"pay_later_name,omitempty"`
	PayLaterRedirectUrl             string `json:"pay_later_redirect_url,omitempty"`
	PayNowAssetUrlsDescriptive      string `json:"pay_now_asset_urls_descriptive,omitempty"`
	PayNowAssetUrlsStandard         string `json:"pay_now_asset_urls_standard,omitempty"`
	PayNowName                      string `json:"pay_now_name,omitempty"`
	PayNowRedirectUrl               string `json:"pay_now_redirect_url,omitempty"`
	PayOverTimeAssetUrlsDescriptive string `json:"pay_over_time_asset_urls_descriptive,omitempty"`
	PayOverTimeAssetUrlsStandard    string `json:"pay_over_time_asset_urls_standard,omitempty"`
	PayOverTimeName                 string `json:"pay_over_time_name,omitempty"`
	PayOverTimeRedirectUrl          string `json:"pay_over_time_redirect_url,omitempty"`
	PaymentMethodCategories         string `json:"payment_method_categories,omitempty"`
	PurchaseCountry                 string `json:"purchase_country,omitempty"`
	PurchaseType                    string `json:"purchase_type,omitempty"`
	RedirectUrl                     string `json:"redirect_url,omitempty"`
	ShippingDelay                   int    `json:"shipping_delay,omitempty"`
	ShippingFirstName               string `json:"shipping_first_name,omitempty"`
	ShippingLastName                string `json:"shipping_last_name,omitempty"`
}
