package intent

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/internal/mandate"
	"github.com/driver005/gateway/payment/method"
	"github.com/driver005/gateway/utils/errors"
	"github.com/google/uuid"
)

type SetupIntentNextActionVerifyWithMicrodeposits struct {
	core.Model

	ArrivalDate           int    `json:"arrival_date,omitempty"`
	HostedVerificationUrl string `json:"hosted_verification_url,omitempty"`
	MicrodepositType      string `json:"microdeposit_type,omitempty"`
}

type SetupIntentNextAction struct {
	core.Model

	RedirectToUrlReturnUrl  string                                        `json:"redirect_to_url_return_url,omitempty"`
	RedirectToUrlUrl        string                                        `json:"redirect_to_url_url,omitempty"`
	Type                    string                                        `json:"type,omitempty"`
	UseStripeSdk            core.JSONB                                    `json:"use_stripe_sdk,omitempty"`
	VerifyWithMicrodeposits *SetupIntentNextActionVerifyWithMicrodeposits `json:"verify_with_microdeposits,omitempty" database:"foreignKey:id"`
}

type SetupIntent struct {
	core.Model

	AttachToSelf       bool                  `json:"attach_to_self,omitempty"`
	CancellationReason string                `json:"cancellation_reason,omitempty"`
	ClientSecret       string                `json:"client_secret,omitempty"`
	Description        string                `json:"description,omitempty"`
	FlowDirections     []string              `json:"flow_directions,omitempty" database:"type:text[]"`
	LastSetupError     errors.ApiErrors      `json:"last_setup_error,omitempty" database:"foreignKey:id"`
	NextAction         SetupIntentNextAction `json:"next_action,omitempty" database:"foreignKey:id"`
	PaymentMethodTypes []string              `json:"payment_method_types" database:"type:text[]"`
	Status             string                `json:"status,omitempty"`
	Usage              string                `json:"usage,omitempty"`

	CustomerId         uuid.UUID             `json:"customer_id"`
	Customer           *customer.Customer    `json:"customer,omitempty" database:"foreignKey:customer_id"`
	PaymentMethodId    uuid.UUID             `json:"payment_method_id"`
	PaymentMethod      *method.PaymentMethod `json:"payment_method,omitempty" database:"foreignKey:payment_method_id"`
	MandateId          uuid.UUID             `json:"mandate_id,omitempty"`
	Mandate            *mandate.Mandate      `json:"mandate,omitempty" database:"foreignKey:mandate_id"`
	SingleUseMandateId uuid.UUID             `json:"single_use_mandate_id,omitempty"`
	SingleUseMandate   *mandate.Mandate      `json:"single_use_mandate,omitempty" database:"foreignKey:single_use_mandate_id"`
}
