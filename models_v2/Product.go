package models

import (
	"github.com/driver005/gateway/core"
	"github.com/google/uuid"
)

type Product struct {
	core.Model

	// A product collection object. Available if the relation `collection` is expanded.
	Collection *ProductCollection `json:"collection" database:"foreignKey:id;references:collection_id"`

	// The Product Collection that the Product belongs to
	CollectionId uuid.NullUUID `json:"collection_id,omitempty"`

	// A short description of the Product.
	Description string `json:"description" database:"default:null"`

	// Whether the Product can be discounted. Discounts will not apply to Line Items of this Product when this flag is set to `false`.
	Discountable bool `json:"discountable" database:"default:false"`

	// The external ID of the product
	ExternalId string `json:"external_id,omitempty"`

	// A unique identifier for the Product (e.g. for slug structure).
	Handle string `json:"handle" database:"default:null"`

	// The height of the Product Variant. May be used in shipping rate calculations.
	Height float64 `json:"height" database:"default:null"`

	// The Harmonized System code of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	HsCode string `json:"hs_code" database:"default:null"`

	// Images of the Product. Available if the relation `images` is expanded.
	Images []Image `json:"images" database:"foreignKey:id"`

	// Whether the Product represents a Gift Card. Products that represent Gift Cards will automatically generate a redeemable Gift Card code once they are purchased.
	IsGiftcard bool `json:"is_giftcard" database:"default:false"`

	// The length of the Product Variant. May be used in shipping rate calculations.
	Length float64 `json:"length" database:"default:null"`

	// The material and composition that the Product Variant is made of, May be used by Fulfillment Providers to pass customs information to shipping carriers.
	Material string `json:"material" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`

	// The Manufacturers Identification code that identifies the manufacturer of the Product Variant. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	MidCode string `json:"mid_code" database:"default:null"`

	// The Product Options that are defined for the Product. Product Variants of the Product will have a unique combination of Product Option Values. Available if the relation `options` is expanded.
	Options []ProductOption `json:"options" database:"foreignKey:id"`

	// The country in which the Product Variant was produced. May be used by Fulfillment Providers to pass customs information to shipping carriers.
	OriginCountry string `json:"origin_country" database:"default:null"`

	// Shipping Profiles have a set of defined Shipping Options that can be used to fulfill a given set of Products.
	Profile *ShippingProfile `json:"profile" database:"foreignKey:id;references:profile_id"`

	// The ID of the Shipping Profile that the Product belongs to. Shipping Profiles have a set of defined Shipping Options that can be used to Fulfill a given set of Products.
	ProfileId uuid.NullUUID `json:"profile_id"`

	// The sales channels the product is associated with. Available if the relation `sales_channels` is expanded.
	SalesChannels *SalesChannel `json:"sales_channels" database:"foreignKey:id"`

	// The status of the product
	Status ProductStatus `json:"status" database:"default:draft"`

	// An optional subtitle that can be used to further specify the Product.
	Subtitle string `json:"subtitle" database:"default:null"`

	// The Product Tags assigned to the Product. Available if the relation `tags` is expanded.
	Tags []ProductTag `json:"tags" database:"foreignKey:id"`

	// A URL to an image file that can be used to identify the Product.
	Thumbnail string `json:"thumbnail" database:"default:null"`

	// A title that can be displayed for easy identification of the Product.
	Title string `json:"title"`

	// Product Type can be added to Products for filtering and reporting purposes.
	Type *ProductType `json:"type" database:"foreignKey:id;references:type_id"`

	// The Product type that the Product belongs to
	TypeId uuid.NullUUID `json:"type_id,omitempty"`

	// The Product Variants that belong to the Product. Each will have a unique combination of Product Option Values. Available if the relation `variants` is expanded.
	Variants []ProductVariant `json:"variants" database:"foreignKey:id"`

	// The weight of the Product Variant. May be used in shipping rate calculations.
	Weight float64 `json:"weight" database:"default:null"`

	// The width of the Product Variant. May be used in shipping rate calculations.
	Width float64 `json:"width" database:"default:null"`
}

// func (p *Product) BeforeCreate(tx *database.DB) (err error) {
// 	p.Id, err = uuid.NewUUID()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
