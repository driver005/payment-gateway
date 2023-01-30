package attempt

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/method/details"
	"github.com/driver005/gateway/utils/errors"
	"github.com/google/uuid"
)

type SetupAttempt struct {
	core.Model

	AttachToSelf         bool                         `json:"attach_to_self,omitempty"`
	FlowDirections       []string                     `json:"flow_directions,omitempty" database:"type:text[]"`
	PaymentMethodDetails details.PaymentMethodDetails `json:"payment_method_details,omitempty" database:"foreignKey:id"`
	SetupError           errors.ApiErrors             `json:"setup_error,omitempty" database:"foreignKey:id"`
	Status               string                       `json:"status,omitempty"`
	Usage                string                       `json:"usage,omitempty"`

	CustomerId      uuid.UUID             `json:"customer_id"`
	Customer        *customer.Customer    `json:"customer,omitempty" database:"foreignKey:customer_id"`
	PaymentMethodId uuid.UUID             `json:"payment_method_id"`
	PaymentMethod   *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id"`
	SetupIntentId   uuid.UUID             `json:"setup_intent_id,omitempty"`
	SetupIntent     *intent.SetupIntent   `json:"setup_intent,omitempty" database:"foreignKey:setup_intent_id"`
}
