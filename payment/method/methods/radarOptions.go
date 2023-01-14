package methods

// PaymentMethodRadarOptions
type PaymentMethodRadarOptions struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `json:"session,omitempty"`
}
