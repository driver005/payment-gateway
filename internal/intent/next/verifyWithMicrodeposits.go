package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionVerifyWithMicrodeposits
type PaymentIntentNextActionVerifyWithMicrodeposits struct {
	core.Model

	// The timestamp when the microdeposits are expected to land.
	ArrivalDate int `json:"arrival_date,omitempty"`
	// The URL for the hosted verification page, which allows customers to verify their bank account.
	HostedVerificationUrl string `json:"hosted_verification_url,omitempty"`
	// The type of the microdeposit sent to the customer. Used to distinguish between different verification methods.
	MicrodepositType string `json:"microdeposit_type,omitempty"`
}
