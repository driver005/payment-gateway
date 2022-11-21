/*
 * Medusa Admin API
 *
 * API reference for Medusa's Admin endpoints. All endpoints are prefixed with `/admin`.  ## Authentication  There are two ways to send authenticated requests to the Medusa server: Using a user's API token, or using a Cookie Session ID.  <!-- ReDoc-Inject: <SecurityDefinitions> --> 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

type PostRegionsRegionRequest struct {

	// The name of the Region
	Name string `json:"name,omitempty"`

	// The 3 character ISO currency code to use for the Region.
	CurrencyCode string `json:"currency_code,omitempty"`

	// If true Medusa will automatically calculate taxes for carts in this region. If false you have to manually call POST /carts/:id/taxes.
	AutomaticTaxes bool `json:"automatic_taxes,omitempty"`

	// Whether gift cards in this region should be applied sales tax when purchasing a gift card
	GiftCardsTaxable bool `json:"gift_cards_taxable,omitempty"`

	// The ID of the tax provider to use; if null the system tax provider is used
	TaxProviderId string `json:"tax_provider_id,omitempty"`

	// An optional tax code the Region.
	TaxCode string `json:"tax_code,omitempty"`

	// The tax rate to use on Orders in the Region.
	TaxRate float32 `json:"tax_rate,omitempty"`

	// [EXPERIMENTAL] Tax included in prices of region
	IncludesTax bool `json:"includes_tax,omitempty"`

	// A list of Payment Provider IDs that should be enabled for the Region
	PaymentProviders []string `json:"payment_providers,omitempty"`

	// A list of Fulfillment Provider IDs that should be enabled for the Region
	FulfillmentProviders []string `json:"fulfillment_providers,omitempty"`

	// A list of countries' 2 ISO Characters that should be included in the Region.
	Countries []string `json:"countries,omitempty"`
}
