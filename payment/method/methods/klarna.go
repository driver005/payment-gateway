package methods

import "github.com/driver005/gateway/core"

type PaymentMethodKlarna struct {
	core.Model

	Day   int `json:"day,omitempty"`
	Month int `json:"month,omitempty"`
	Year  int `json:"year,omitempty"`
}
