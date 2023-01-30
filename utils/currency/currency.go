package currency

import (
	"github.com/driver005/gateway/core"
)

type Currency struct {
	core.Model

	// The 3 character ISO code for the currency.
	Code string `json:"code" database:"primarykey"`

	// The symbol used to indicate the currency.
	Symbol string `json:"symbol"`

	// The native symbol used to indicate the currency.
	SymbolNative string `json:"symbol_native"`

	// The written name of the currency
	Name string `json:"name"`

	// [EXPERIMENTAL] Does the currency prices include tax
	IncludesTax bool `json:"includes_tax" database:"default:null"`
}
