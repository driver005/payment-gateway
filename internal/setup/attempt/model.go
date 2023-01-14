package attempt

import (
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/models"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/method/details"
)

// SetupAttempt A SetupAttempt describes one attempted confirmation of a SetupIntent, whether that confirmation was successful or unsuccessful. You can use SetupAttempts to inspect details of a specific attempt at setting up a payment method using a SetupIntent.
type SetupAttempt struct {
	// If present, the SetupIntent's payment method will be attached to the in-context Stripe Account.  It can only be used for this Stripe Account’s own money movement flows like InboundTransfer and OutboundTransfers. It cannot be set to true when setting up a PaymentMethod for a Customer, and defaults to false when attaching a PaymentMethod to a Customer.
	AttachToSelf *bool `json:"attach_to_self,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created  int             `json:"created"`
	Customer customer.Customer `json:"customer,omitempty"`
	// Indicates the directions of money movement for which this payment method is intended to be used.  Include `inbound` if you intend to use the payment method as the origin to pull funds from. Include `outbound` if you intend to use the payment method as the destination to send funds to. You can include both if you intend to use the payment method for both purposes.
	FlowDirections []string `json:"flow_directions,omitempty"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object               string                       `json:"object"`
	PaymentMethod        method.PaymentMethod         `json:"payment_method"`
	PaymentMethodDetails details.PaymentMethodDetails `json:"payment_method_details"`
	SetupError           models.ApiErrors             `json:"setup_error,omitempty"`
	SetupIntent          intent.SetupIntent           `json:"setup_intent"`
	// Status of this SetupAttempt, one of `requires_confirmation`, `requires_action`, `processing`, `succeeded`, `failed`, or `abandoned`.
	Status string `json:"status"`
	// The value of [usage](https://stripe.com/docs/api/setup_intents/object#setup_intent_object-usage) on the SetupIntent at the time of this confirmation, one of `off_session` or `on_session`.
	Usage string `json:"usage"`
}
