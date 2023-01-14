package next

// PaymentIntentNextActionPromptpayDisplayQrCode
type PaymentIntentNextActionPromptpayDisplayQrCode struct {
	// The raw data string used to generate QR code, it should be used together with QR code library.
	Data string `json:"data"`
	// The URL to the hosted PromptPay instructions page, which allows customers to view the PromptPay QR code.
	HostedInstructionsUrl string `json:"hosted_instructions_url"`
	// The PNG path used to render the QR code, can be used as the source in an HTML img tag
	ImageUrlPng string `json:"image_url_png"`
	// The SVG path used to render the QR code, can be used as the source in an HTML img tag
	ImageUrlSvg string `json:"image_url_svg"`
}
