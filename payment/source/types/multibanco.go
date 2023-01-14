package types

// SourceTypeMultibanco struct for SourceTypeMultibanco
type SourceTypeMultibanco struct {
	Entity string `json:"entity,omitempty"`
	Reference string `json:"reference,omitempty"`
	RefundAccountHolderAddressCity string `json:"refund_account_holder_address_city,omitempty"`
	RefundAccountHolderAddressCountry string `json:"refund_account_holder_address_country,omitempty"`
	RefundAccountHolderAddressLine1 string `json:"refund_account_holder_address_line1,omitempty"`
	RefundAccountHolderAddressLine2 string `json:"refund_account_holder_address_line2,omitempty"`
	RefundAccountHolderAddressPostalCode string `json:"refund_account_holder_address_postal_code,omitempty"`
	RefundAccountHolderAddressState string `json:"refund_account_holder_address_state,omitempty"`
	RefundAccountHolderName string `json:"refund_account_holder_name,omitempty"`
	RefundIban string `json:"refund_iban,omitempty"`
}