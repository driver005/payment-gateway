package types

import (
	"github.com/driver005/gateway/core"
	"github.com/gofrs/uuid"
)

type FilterableLineItemProps struct {
	core.Filter
	ItemId      []uuid.NullUUID `json:"item_id,omitempty"`
	Description []string        `json:"description,omitempty"`
	ResourceId  []uuid.NullUUID `json:"resource_id,omitempty"`
}
