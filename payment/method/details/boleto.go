package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsBoleto
type PaymentMethodDetailsBoleto struct {
	core.Model

	// The tax ID of the customer (CPF for individuals consumers or CNPJ for businesses consumers)
	TaxId string `json:"tax_id,omitempty"`
}
