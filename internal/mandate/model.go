package mandate

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/mandate/methods"
	"github.com/driver005/gateway/payment/method"
	"github.com/google/uuid"
)

type Mandate struct {
	core.Model

	PaymentMethodDetails *methods.MandatePaymentMethodDetails `json:"payment_method_details,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SingleUse            *methods.MandateSingleUse            `json:"single_use,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Status               Status                               `json:"status,omitempty"`
	Type                 Type                                 `json:"type,omitempty"`

	PaymentMethodId *uuid.UUID            `json:"payment_method_id,omitempty" swaggerignore:"true"`
	PaymentMethod   *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id" swaggertype:"primitive,string" format:"uuid"`
}
