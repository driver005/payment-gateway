package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionPaynowDisplayQrCode
type PaymentIntentNextActionPaynowDisplayQrCode struct {
	core.Model

	// The raw data string used to generate QR code, it should be used together with QR code library.
	Data string `json:"data,omitempty"`
	// The URL to the hosted PayNow instructions page, which allows customers to view the PayNow QR code.
	HostedInstructionsUrl string `json:"hosted_instructions_url,omitempty"`
	// The image_url_png string used to render QR code
	ImageUrlPng string `json:"image_url_png,omitempty"`
	// The image_url_svg string used to render QR code
	ImageUrlSvg string `json:"image_url_svg,omitempty"`
}
