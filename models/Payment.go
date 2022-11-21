package models

import (
	"time"

	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

// Payment - Payments represent an amount authorized with a given payment method, Payments can be captured, canceled or refunded.
type Payment struct {
	core.Model

	// The ID of the Swap that the Payment is used for.
	SwapId uuid.NullUUID `json:"swap_id" database:"default:null"`

	// A swap object. Available if the relation `swap` is expanded.
	Swap *Swap `json:"swap" database:"foreignKey:id;references:swap_id"`

	// The id of the Cart that the Payment Session is created for.
	CartId uuid.NullUUID `json:"cart_id" database:"default:null"`

	// A cart object. Available if the relation `cart` is expanded.
	Cart *Cart `json:"cart" database:"foreignKey:id;references:cart_id"`

	// The ID of the Order that the Payment is used for.
	OrderId uuid.NullUUID `json:"order_id" database:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// The amount that the Payment has been authorized for.
	Amount int32 `json:"amount"`

	// The 3 character ISO currency code that the Payment is completed in.
	CurrencyCode string `json:"currency_code"`

	Currency *Currency `json:"currency" database:"foreignKey:code;references:currency_code"`

	// The amount of the original Payment amount that has been refunded back to the Customer.
	AmountRefunded int32 `json:"amount_refunded" database:"default:null"`

	// The id of the Payment Provider that is responsible for the Payment
	ProviderId uuid.NullUUID `json:"provider_id"`

	// The data required for the Payment Provider to identify, modify and process the Payment. Typically this will be an object that holds an id to the external payment session, but can be an empty object if the Payment Provider doesn't hold any state.
	Data JSONB `json:"data" database:"default:null"`

	// The date with timezone at which the Payment was captured.
	CapturedAt time.Time `json:"captured_at" database:"default:null"`

	// The date with timezone at which the Payment was canceled.
	CanceledAt time.Time `json:"canceled_at" database:"default:null"`

	// Randomly generated key used to continue the completion of a payment in case of failure.
	IdempotencyKey string `json:"idempotency_key" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
