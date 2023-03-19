package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// NotificationResend - A resend of a Notification.
type NotificationResend struct {
	core.Model

	// The name of the event that the notification was sent for.
	EventName string `json:"event_name" database:"default:null"`

	// The type of resource that the Notification refers to.
	ResourceType string `json:"resource_type" database:"default:null"`

	// The ID of the resource that the Notification refers to.
	ResourceId uuid.NullUUID `json:"resource_id,omitempty"`

	// The ID of the Customer that the Notification was sent to.
	CustomerId uuid.NullUUID `json:"customer_id,omitempty"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" database:"foreignKey:id;references:customer_id"`

	// The address that the Notification was sent to. This will usually be an email address, but represent other addresses such as a chat bot user id
	To string `json:"to" database:"default:null"`

	// The data that the Notification was sent with. This contains all the data necessary for the Notification Provider to initiate a resend.
	Data JSONB `json:"data" database:"default:null"`

	// The ID of the Notification that was originally sent.
	ParentId uuid.NullUUID `json:"parent_id,omitempty"`

	ParentNotification *Notification `json:"parent_notification" database:"foreignKey:id;references:parent_id"`

	// The ID of the Notification Provider that handles the Notification.
	ProviderId uuid.NullUUID `json:"provider_id,omitempty"`

	Provider *NotificationProvider `json:"provider" database:"foreignKey:id;references:provider_id"`
}
