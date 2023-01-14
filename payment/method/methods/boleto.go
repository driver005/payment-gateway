package methods

// PaymentMethodBoleto
type PaymentMethodBoleto struct {
	// Uniquely identifies the customer tax id (CNPJ or CPF)
	TaxId string `json:"tax_id"`
}
