package next

// PaymentIntentNextActionBoleto
type PaymentIntentNextActionBoleto struct {
	// The timestamp after which the boleto expires.
	ExpiresAt int `json:"expires_at,omitempty"`
	// The URL to the hosted boleto voucher page, which allows customers to view the boleto voucher.
	HostedVoucherUrl string `json:"hosted_voucher_url,omitempty"`
	// The boleto number.
	Number string `json:"number,omitempty"`
	// The URL to the downloadable boleto voucher PDF.
	Pdf string `json:"pdf,omitempty"`
}
