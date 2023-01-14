package models

// ThreeDSecureDetails
type ThreeDSecureDetails struct {
	// For authenticated transactions: how the customer was authenticated by the issuing bank.
	AuthenticationFlow string `json:"authentication_flow,omitempty"`
	// Indicates the outcome of 3D Secure authentication.
	Result string `json:"result,omitempty"`
	// Additional information about why 3D Secure succeeded or failed based on the `result`.
	ResultReason string `json:"result_reason,omitempty"`
	// The version of 3D Secure that was used.
	Version string `json:"version,omitempty"`
}
