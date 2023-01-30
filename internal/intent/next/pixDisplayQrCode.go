package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionPixDisplayQrCode
type PaymentIntentNextActionPixDisplayQrCode struct {
	core.Model

	// The raw data string used to generate QR code, it should be used together with QR code library.
	Data string `json:"data,omitempty"`
	// The date (unix timestamp) when the PIX expires.
	ExpiresAt int `json:"expires_at,omitempty"`
	// The URL to the hosted pix instructions page, which allows customers to view the pix QR code.
	HostedInstructionsUrl string `json:"hosted_instructions_url,omitempty"`
	// The image_url_png string used to render png QR code
	ImageUrlPng string `json:"image_url_png,omitempty"`
	// The image_url_svg string used to render svg QR code
	ImageUrlSvg string `json:"image_url_svg,omitempty"`
}
