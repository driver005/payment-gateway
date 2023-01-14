package next

// PaymentIntentNextActionVerifyWithMicrodeposits
type PaymentIntentNextActionVerifyWithMicrodeposits struct {
	// The timestamp when the microdeposits are expected to land.
	ArrivalDate int `json:"arrival_date"`
	// The URL for the hosted verification page, which allows customers to verify their bank account.
	HostedVerificationUrl string `json:"hosted_verification_url"`
	// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
	MicrodepositType string `json:"microdeposit_type,omitempty"`
}
