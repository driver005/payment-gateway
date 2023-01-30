package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionKonbini
type PaymentIntentNextActionKonbini struct {
	core.Model

	// The timestamp at which the pending Konbini payment expires.
	ExpiresAt int `json:"expires_at"`
	// The URL for the Konbini payment instructions page, which allows customers to view and print a Konbini voucher.
	HostedVoucherUrl string                               `json:"hosted_voucher_url,omitempty"`
	Stores           PaymentIntentNextActionKonbiniStores `json:"stores" database:"foreignKey:id"`
}
