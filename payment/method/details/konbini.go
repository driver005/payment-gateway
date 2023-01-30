package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsKonbini
type PaymentMethodDetailsKonbini struct {
	core.Model

	StoreChain string `json:"store_chain,omitempty"`
}
