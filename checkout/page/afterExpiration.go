package page

import "github.com/driver005/gateway/core"

// PaymentPagesCheckoutSessionAfterExpiration
type PaymentPagesCheckoutSessionAfterExpiration struct {
	core.Model

	// Enables user redeemable promotion codes on the recovered Checkout Sessions. Defaults to `false`
	AllowPromotionCodes bool `json:"allow_promotion_codes,omitempty"`
	// If `true`, a recovery url will be generated to recover this Checkout Session if it expires before a transaction is completed. It will be attached to the Checkout Session object upon expiration.
	Enabled bool `json:"enabled,omitempty"`
	// The timestamp at which the recovery URL will expire.
	ExpiresAt int `json:"expires_at,omitempty"`
	// URL that creates a new Checkout Session when clicked that is a copy of this expired Checkout Session
	Url string `json:"url,omitempty"`
}
