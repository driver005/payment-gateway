package event

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/account"
)

type NotificationEventData struct {
	core.Model

	Object             core.JSONB `json:"object,omitempty"`
	PreviousAttributes core.JSONB `json:"previous_attributes,omitempty"`
}

type NotificationEventRequest struct {
	core.Model

	Id             string `json:"id,omitempty"`
	IdempotencyKey string `json:"idempotency_key,omitempty"`
}

type Event struct {
	core.Model

	Account         *account.Account          `json:"account,omitempty" database:"foreignKey:id"`
	ApiVersion      string                    `json:"api_version,omitempty"`
	Data            *NotificationEventData    `json:"data" database:"foreignKey:id"`
	PendingWebhooks int                       `json:"pending_webhooks,omitempty"`
	Request         *NotificationEventRequest `json:"request,omitempty" database:"foreignKey:id"`
	Type            string                    `json:"type,omitempty"`
}
