package mandate

import (
	"github.com/driver005/gateway/internal/mandate/methods"
	"github.com/driver005/gateway/payment/method"
)

type Mandate struct {
	Id string `json:"id"`
	Livemode bool `json:"livemode"`
	MultiUse map[string]interface{} `json:"multi_use,omitempty"`
	Object               string                      `json:"object"`
	PaymentMethod        *method.PaymentMethod        `json:"payment_method"`
	PaymentMethodDetails *methods.MandatePaymentMethodDetails `json:"payment_method_details"`
	SingleUse            *methods.MandateSingleUse           `json:"single_use,omitempty"`
	Status string `json:"status"`
	Type string `json:"type"`
}
