package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionDisplayOxxoDetails
type PaymentIntentNextActionDisplayOxxoDetails struct {
	core.Model

	// The timestamp after which the OXXO voucher expires.
	ExpiresAfter int `json:"expires_after,omitempty"`
	// The URL for the hosted OXXO voucher page, which allows customers to view and print an OXXO voucher.
	HostedVoucherUrl string `json:"hosted_voucher_url,omitempty"`
	// OXXO reference number.
	Number string `json:"number,omitempty"`
}
