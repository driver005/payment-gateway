package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Notification - Notifications a communications sent via Notification Providers as a reaction to internal events such as `order.placed`. Notifications can be used to show a chronological timeline for communications sent to a Customer regarding an Order, and enables resends.
type Notification struct {
	core.Model

	// The name of the event that the notification was sent for.
	EventName string `json:"event_name" database:"default:null"`

	// The type of resource that the Notification refers to.
	ResourceType string `json:"resource_type"`

	// The ID of the resource that the Notification refers to.
	ResourceId uuid.NullUUID `json:"resource_id"`

	// The ID of the Customer that the Notification was sent to.
	CustomerId uuid.NullUUID `json:"customer_id" database:"default:null"`

	// A customer object. Available if the relation `customer` is expanded.
	Customer *Customer `json:"customer" database:"foreignKey:id;references:customer_id"`

	// The address that the Notification was sent to. This will usually be an email address, but represent other addresses such as a chat bot user id
	To string `json:"to"`

	// The data that the Notification was sent with. This contains all the data necessary for the Notification Provider to initiate a resend.
	Data JSONB `json:"data" database:"default:null"`

	// The resends that have been completed after the original Notification.
	Resends []NotificationResend `json:"resends" database:"foreignKey:id"`

	// The id of the Notification Provider that handles the Notification.
	ProviderId uuid.NullUUID `json:"provider_id" database:"default:null"`

	Provider *NotificationProvider `json:"provider" database:"foreignKey:id;references:provider_id"`
}
