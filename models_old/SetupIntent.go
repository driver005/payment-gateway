package models

type SetupIntent struct {
	ID                 string      `json:"id"`
	Object             string      `json:"object"`
	Application        interface{} `json:"application"`
	CancellationReason interface{} `json:"cancellation_reason"`
	ClientSecret       string      `json:"client_secret"`
	Created            int         `json:"created"`
	Customer           interface{} `json:"customer"`
	Description        interface{} `json:"description"`
	LastSetupError     interface{} `json:"last_setup_error"`
	LatestAttempt      interface{} `json:"latest_attempt"`
	Livemode           bool        `json:"livemode"`
	Mandate            interface{} `json:"mandate"`
	Metadata           struct {
	} `json:"metadata"`
	NextAction           interface{} `json:"next_action"`
	OnBehalfOf           interface{} `json:"on_behalf_of"`
	PaymentMethod        interface{} `json:"payment_method"`
	PaymentMethodOptions struct {
		Card struct {
			MandateOptions      interface{} `json:"mandate_options"`
			RequestThreeDSecure string      `json:"request_three_d_secure"`
		} `json:"card"`
	} `json:"payment_method_options"`
	PaymentMethodTypes []string    `json:"payment_method_types"`
	SingleUseMandate   interface{} `json:"single_use_mandate"`
	Status             string      `json:"status"`
	Usage              string      `json:"usage"`
}
