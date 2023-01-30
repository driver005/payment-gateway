package radar

import (
	"github.com/driver005/gateway/core"
)

// RadarRadarOptions Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type RadarOptions struct {
	core.Model
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session string `json:"session,omitempty"`
}
