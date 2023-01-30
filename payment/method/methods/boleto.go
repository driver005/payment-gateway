package methods

import "github.com/driver005/gateway/core"

type PaymentMethodBoleto struct {
	core.Model

	TaxId string `json:"tax_id,omitempty"`
}
