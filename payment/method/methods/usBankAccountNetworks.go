package methods

// PaymentMethodUsBankAccountNetworks Contains information about US bank account networks that can be used.
type PaymentMethodUsBankAccountNetworks struct {
	// The preferred network.
	Preferred string `json:"preferred,omitempty"`
	// All supported networks.
	Supported []string `json:"supported"`
}
