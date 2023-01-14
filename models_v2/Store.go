package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

// Store - Holds settings for the Store, such as name, currencies, etc.
type Store struct {
	core.Model

	// The name of the Store - this may be displayed to the Customer.
	Name string `json:"name" database:"default:null"`

	// The 3 character currency code that is the default of the store.
	DefaultCurrencyCode string `json:"default_currency_code" database:"default:null"`

	DefaultCurrency *Currency `json:"default_currency" database:"foreignKey:id;references:default_currency_code"`

	// The currencies that are enabled for the Store. Available if the relation `currencies` is expanded.
	Currencies []Currency `json:"currencies" database:"foreignKey:id"`

	// A template to generate Swap links from. Use {{cart_id}} to include the Swap's `cart_id` in the link.
	SwapLinkTemplate string `json:"swap_link_template" database:"default:null"`

	// A template to generate Payment links from. Use {{cart_id}} to include the payment's `cart_id` in the link.
	PaymentLinkTemplate string `json:"payment_link_template" database:"default:null"`

	// A template to generate Invite links from
	InviteLinkTemplate string `json:"invite_link_template" database:"default:null"`

	// The sales channel ID the cart is associated with.
	DefaultSalesChannelId uuid.NullUUID `json:"default_sales_channel_id" database:"default:null"`

	// A sales channel object. Available if the relation `default_sales_channel` is expanded.
	DefaultSalesChannel *SalesChannel `json:"default_sales_channel" database:"foreignKey:id;references:default_sales_channel_id"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
