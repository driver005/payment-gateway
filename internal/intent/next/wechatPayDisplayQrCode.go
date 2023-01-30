package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionWechatPayDisplayQrCode
type PaymentIntentNextActionWechatPayDisplayQrCode struct {
	core.Model

	// The data being used to generate QR code
	Data string `json:"data,omitempty"`
	// The URL to the hosted WeChat Pay instructions page, which allows customers to view the WeChat Pay QR code.
	HostedInstructionsUrl string `json:"hosted_instructions_url,omitempty"`
	// The base64 image data for a pre-generated QR code
	ImageDataUrl string `json:"image_data_url,omitempty"`
	// The image_url_png string used to render QR code
	ImageUrlPng string `json:"image_url_png,omitempty"`
	// The image_url_svg string used to render QR code
	ImageUrlSvg string `json:"image_url_svg,omitempty"`
}
