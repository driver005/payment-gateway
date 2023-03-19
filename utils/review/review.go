package review

import (
	"github.com/driver005/gateway/core"
)

// RadarReviewResourceSession
type RadarReviewResourceSession struct {
	core.Model

	// The browser used in this browser session (e.g., `Chrome`).
	Browser string `json:"browser,omitempty"`
	// Information about the device used for the browser session (e.g., `Samsung SM-G930T`).
	Device string `json:"device,omitempty"`
	// The platform for the browser session (e.g., `Macintosh`).
	Platform string `json:"platform,omitempty"`
	// The version for the browser session (e.g., `61.0.3163.100`).
	Version string `json:"version,omitempty"`
}

// RadarReviewResourceLocation
type RadarReviewResourceLocation struct {
	core.Model

	// The city where the payment originated.
	City string `json:"city,omitempty"`
	// Two-letter ISO code representing the country where the payment originated.
	Country string `json:"country,omitempty"`
	// The geographic latitude where the payment originated.
	Latitude float64 `json:"latitude,omitempty"`
	// The geographic longitude where the payment originated.
	Longitude float64 `json:"longitude,omitempty"`
	// The state/county/province/region where the payment originated.
	Region string `json:"region,omitempty"`
}

// Review Reviews can be used to supplement automated fraud detection with human expertise.  Learn more about [Radar](/radar) and reviewing payments [here](https://stripe.com/docs/radar/reviews).
type Review struct {
	core.Model

	// The ZIP or postal code of the card used, if applicable.
	BillingZip string `json:"billing_zip,omitempty"`
	// The reason the review was closed, or null if it has not yet been closed. One of `approved`, `refunded`, `refunded_as_fraud`, `disputed`, or `redacted`.
	ClosedReason string `json:"closed_reason,omitempty"`
	// The IP address where the payment originated.
	IpAddress         string                       `json:"ip_address,omitempty"`
	IpAddressLocation *RadarReviewResourceLocation `json:"ip_address_location,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// If `true`, the review needs action.
	Open bool `json:"open,omitempty"`
	// The reason the review was opened. One of `rule` or `manual`.
	OpenedReason string `json:"opened_reason,omitempty"`
	// The reason the review is currently open or closed. One of `rule`, `manual`, `approved`, `refunded`, `refunded_as_fraud`, `disputed`, or `redacted`.
	Reason  string                      `json:"reason,omitempty"`
	Session *RadarReviewResourceSession `json:"session,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
