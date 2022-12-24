package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/google/uuid"
)

type FilterableProductProps struct {
	core.Filter
	Q                   string               `json:"q,omitempty"`
	Status              models.ProductStatus `json:"status,omitempty"`
	PriceListId         []uuid.NullUUID      `json:"price_list_id,omitempty"`
	CollectionId        []uuid.NullUUID      `json:"collection_id,omitempty"`
	Tags                []string             `json:"tags,omitempty"`
	Title               string               `json:"title,omitempty"`
	Description         string               `json:"description,omitempty"`
	Handle              string               `json:"handle,omitempty"`
	IsGiftcard          bool                 `json:"is_giftcard,omitempty"`
	TypeId              []uuid.NullUUID      `json:"type_id,omitempty"`
	SalesChannelId      []uuid.NullUUID      `json:"sales_channel_id,omitempty"`
	DiscountConditionId []uuid.NullUUID      `json:"discount_condition_id,omitempty"`
}
