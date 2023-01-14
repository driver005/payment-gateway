package methods

// PaymentMethodIdeal
type PaymentMethodIdeal struct {
	// The customer's bank, if provided. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
	Bank string `json:"bank,omitempty"`
	// The Bank Identifier Code of the customer's bank, if the bank was provided.
	Bic string `json:"bic,omitempty"`
}
