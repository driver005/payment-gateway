package types

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/models"
	"github.com/gofrs/uuid"
)

type FilterablePriceListProps struct {
	core.Filter
	Q              string                   `json:"q,omitempty"`
	Status         []models.PriceListStatus `json:"status,omitempty"`
	Name           string                   `json:"name,omitempty"`
	CustomerGroups []string                 `json:"customer_groups,omitempty"`
	Type           []models.PriceListType   `json:"type,omitempty"`
	ResourceId     uuid.NullUUID            `json:"resource_id,omitempty"`
}
