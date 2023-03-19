package attempt

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/setup/intent"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/payment/method/details"
	"github.com/driver005/gateway/utils/errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SetupAttempt struct {
	core.Model

	AttachToSelf         bool                          `json:"attach_to_self,omitempty"`
	FlowDirections       pq.StringArray                `json:"flow_directions,omitempty" database:"type:varchar(64)[]"`
	PaymentMethodDetails *details.PaymentMethodDetails `json:"payment_method_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SetupError           *errors.ApiErrors             `json:"setup_error,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Status               Status                        `json:"status,omitempty"`
	Usage                Usage                         `json:"usage,omitempty"`

	CustomerId      *uuid.UUID            `json:"customer_id" swaggerignore:"true"`
	Customer        *customer.Customer    `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
	PaymentMethodId *uuid.UUID            `json:"payment_method_id" swaggerignore:"true"`
	PaymentMethod   *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id" swaggertype:"primitive,string" format:"uuid"`
	SetupIntentId   *uuid.UUID            `json:"setup_intent_id,omitempty" swaggerignore:"true"`
	SetupIntent     *intent.SetupIntent   `json:"setup_intent,omitempty" database:"foreignKey:setup_intent_id" swaggertype:"primitive,string" format:"uuid"`
}
