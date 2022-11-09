package models

import "time"

// Shipping Options represent a way in which an Order or Return can be shipped. Shipping Options have an associated Fulfillment Provider that will be used when the fulfillment of an Order is initiated. Shipping Options themselves cannot be added to Carts, but serve as a template for Shipping Methods. This distinction makes it possible to customize individual Shipping Methods with additional information.
type ShippingOption struct {
	// The amount to charge for shipping when the Shipping Option price type is `flat_rate`.
	Amount *int `json:"amount,omitempty"`

	// The data needed for the Fulfillment Provider to identify the Shipping Option.
	Data *map[string]interface{} `json:"data,omitempty"`

	// The shipping option's ID
	Id *string `json:"id,omitempty"`

	// [EXPERIMENTAL] Does the shipping option price include tax
	IncludesTax *bool `json:"includes_tax,omitempty"`

	// Flag to indicate if the Shipping Option can be used for Return shipments.
	IsReturn *bool `json:"is_return,omitempty"`

	// An optional key-value map with additional details
	Metadata *map[string]interface{} `json:"metadata,omitempty"`

	// The name given to the Shipping Option - this may be displayed to the Customer.
	Name string `json:"name"`

	// The type of pricing calculation that is used when creatin Shipping Methods from the Shipping Option. Can be `flat_rate` for fixed prices or `calculated` if the Fulfillment Provider can provide price calulations.
	PriceType ShippingOptionPriceType `json:"price_type"`

	// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
	Profile *ShippingProfile `json:"profile,omitempty"`

	// The ID of the Shipping Profile that the shipping option belongs to. Shipping Profiles have a set of defined Shipping Options that can be used to Fulfill a given set of Products.
	ProfileId string `json:"profile_id"`

	// Represents a fulfillment provider plugin and holds its installation status.
	Provider *FulfillmentProvider `json:"provider,omitempty"`

	// The id of the Fulfillment Provider, that will be used to process Fulfillments from the Shipping Option.
	ProviderId string `json:"provider_id"`

	// A region object. Available if the relation `region` is expanded.
	Region *map[string]interface{} `json:"region,omitempty"`

	// The region's ID
	RegionId string `json:"region_id"`

	// The requirements that must be satisfied for the Shipping Option to be available for a Cart. Available if the relation `requirements` is expanded.
	Requirements *[]ShippingOptionRequirement `json:"requirements,omitempty"`

	// The date with timezone at which the resource was created.
	CreatedAt *time.Time `json:"created_at,omitempty"`

	// The date with timezone at which the resource was deleted.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// The date with timezone at which the resource was updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
