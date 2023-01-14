package address

// Shipping
type Shipping struct {
	Address *Address `json:"address,omitempty"`
	// The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.
	Carrier string `json:"carrier,omitempty"`
	// Recipient name.
	Name string `json:"name,omitempty"`
	// Recipient phone (including extension).
	Phone string `json:"phone,omitempty"`
	// The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.
	TrackingNumber string `json:"tracking_number,omitempty"`
}
