package options

// InvoicePaymentMethodOptionsCard
type InvoicePaymentMethodOptionsCard struct {
	Installments *InvoiceInstallmentsCard `json:"installments,omitempty"`
	// We strongly recommend that you rely on our SCA Engine to automatically prompt your customers for authentication based on risk level and [other requirements](https://stripe.com/docs/strong-customer-authentication). However, if you wish to request 3D Secure based on logic from your own fraud engine, provide this option. Read our guide on [manually requesting 3D Secure](https://stripe.com/docs/payments/3d-secure#manual-three-ds) for more information on how this configuration interacts with Radar and our SCA Engine.
	RequestThreeDSecure string `json:"request_three_d_secure,omitempty"`
}
