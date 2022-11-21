package models

import (
	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

// GiftCardTransaction - Gift Card Transactions are created once a Customer uses a Gift Card to pay for their Order
type GiftCardTransaction struct {
	core.Model

	// The ID of the Gift Card that was used in the transaction.
	GiftCardId uuid.NullUUID `json:"gift_card_id"`

	// A gift card object. Available if the relation `gift_card` is expanded.
	GiftCard *GiftCard `json:"gift_card" database:"foreignKey:id;references:gift_card_id"`

	// The ID of the Order that the Gift Card was used to pay for.
	OrderId uuid.NullUUID `json:"order_id" database:"default:null"`

	// An order object. Available if the relation `order` is expanded.
	Order *Order `json:"order" database:"foreignKey:id;references:order_id"`

	// The amount that was used from the Gift Card.
	Amount int32 `json:"amount"`

	// Whether the transaction is taxable or not.
	IsTaxable bool `json:"is_taxable" database:"default:null"`

	// The tax rate of the transaction
	TaxRate float32 `json:"tax_rate" database:"default:null"`
}
