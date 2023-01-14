package methods

// MandateBlik
type MandateBlik struct {
	// Date at which the mandate expires.
	ExpiresAfter int                                  `json:"expires_after,omitempty"`
	OffSession   *MandateOptionsOffSessionDetailsBlik `json:"off_session,omitempty"`
	// Type of the mandate.
	Type string `json:"type,omitempty"`
}
