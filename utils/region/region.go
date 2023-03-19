package region

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/currency"
	"github.com/driver005/gateway/utils/tax"
)

// Region - Regions hold settings for how Customers in a given geographical location shop. The is, for example, where currencies and tax rates are defined. A Region can consist of multiple countries to accomodate common shopping settings across countries.
type Region struct {
	core.Model

	// The name of the region as displayed to the customer. If the Region only has one country it is recommended to write the country name.
	Name string `json:"name"`

	// The 3 character currency code that the Region uses.
	CurrencyCode string `json:"currency_code"`

	Currency *currency.Currency `json:"currency" database:"foreignKey:code;references:currency_code"`

	// The tax rate that should be charged on purchases in the Region.
	TaxRate float64 `json:"tax_rate"`

	// The tax rates that are included in the Region. Available if the relation `tax_rates` is expanded.
	TaxRates *tax.TaxRate `json:"tax_rates" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`

	// The tax code used on purchases in the Region. This may be used by other systems for accounting purposes.
	TaxCode string `json:"tax_code" database:"default:null"`

	// Whether the gift cards are taxable or not in this region.
	GiftCardsTaxable bool `json:"gift_cards_taxable" database:"default:null"`

	// Whether taxes should be automated in this region.
	AutomaticTaxes bool `json:"automatic_taxes" database:"default:null"`

	// [EXPERIMENTAL] Does the prices for the region include tax
	IncludesTax bool `json:"includes_tax" database:"default:null"`
}
