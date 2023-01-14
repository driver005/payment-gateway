package methods

// PaymentMethodLink
type PaymentMethodLink struct {
	// Account owner's email address.
	Email string `json:"email,omitempty"`
	// Token used for persistent Link logins.
	PersistentToken string `json:"persistent_token,omitempty"`
}
