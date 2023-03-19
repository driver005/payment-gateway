package intent

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils/errors"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type SetupIntentNextActionVerifyWithMicrodeposits struct {
	core.Model

	ArrivalDate           int              `json:"arrival_date,omitempty"`
	HostedVerificationUrl string           `json:"hosted_verification_url,omitempty"`
	MicrodepositType      MicrodepositType `json:"microdeposit_type,omitempty"`
}

type SetupIntentNextAction struct {
	core.Model

	RedirectToUrlReturnUrl  string                                        `json:"redirect_to_url_return_url,omitempty"`
	RedirectToUrlUrl        string                                        `json:"redirect_to_url_url,omitempty"`
	Type                    Type                                          `json:"type,omitempty"`
	UseStripeSdk            core.JSONB                                    `json:"use_stripe_sdk,omitempty"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

type SetupIntent struct {
	core.Model

	AttachToSelf       bool                   `json:"attach_to_self,omitempty"`
	CancellationReason CancellationReason     `json:"cancellation_reason,omitempty"`
	ClientSecret       string                 `json:"client_secret,omitempty"`
	Description        string                 `json:"description,omitempty"`
	FlowDirections     pq.StringArray         `json:"flow_directions,omitempty" database:"type:varchar(64)[]"`
	LastSetupError     *errors.ApiErrors      `json:"last_setup_error,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	NextAction         *SetupIntentNextAction `json:"next_action,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	PaymentMethodTypes pq.StringArray         `json:"payment_method_types" database:"type:varchar(64)[]"`
	Status             Status                 `json:"status,omitempty"`
	Usage              Usage                  `json:"usage,omitempty"`

	CustomerId         *uuid.UUID            `json:"customer_id" swaggerignore:"true"`
	Customer           *customer.Customer    `json:"customer,omitempty" database:"foreignKey:customer_id" swaggertype:"primitive,string" format:"uuid"`
	PaymentMethodId    *uuid.UUID            `json:"payment_method_id" swaggerignore:"true"`
	PaymentMethod      *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id" swaggertype:"primitive,string" format:"uuid"`
	MandateId          *uuid.UUID            `json:"mandate_id,omitempty" swaggerignore:"true"`
	Mandate            *mandate.Mandate      `json:"mandate,omitempty" database:"foreignKey:mandate_id" swaggertype:"primitive,string" format:"uuid"`
	SingleUseMandateId *uuid.UUID            `json:"single_use_mandate_id,omitempty" swaggerignore:"true"`
	SingleUseMandate   *mandate.Mandate      `json:"single_use_mandate,omitempty" database:"foreignKey:single_use_mandate_id" swaggertype:"primitive,string" format:"uuid"`
}
