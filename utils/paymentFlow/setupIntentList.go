package paymentFlow

import "github.com/driver005/gateway/models"

// PaymentFlowsSetupIntentList
type PaymentFlowsSetupIntentList struct {
	Data []models.SetupIntent `json:"data"`
	// True if this list has another page of items after this one that can be fetched.
	HasMore bool `json:"has_more"`
	// String representing the object's type. Objects of the same type share the same value. Always has the value `list`.
	Object string `json:"object"`
	// The URL where this list can be accessed.
	Url string `json:"url"`
}
