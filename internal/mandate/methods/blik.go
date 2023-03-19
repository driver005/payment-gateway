package methods

import "github.com/driver005/gateway/core"

// MandateBlik
type MandateBlik struct {
	core.Model

	// Date at which the mandate expires.
	ExpiresAfter int                                  `json:"expires_after,omitempty"`
	OffSession   *MandateOptionsOffSessionDetailsBlik `json:"off_session,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// Type of the mandate.
	Type string `json:"type,omitempty"`
}
