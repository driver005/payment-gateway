package core

import (
	"github.com/gofrs/uuid"
)

type Filter struct {
	Id        []uuid.NullUUID `json:"id,omitempty"`
	CreatedAt Time            `json:"created_at,omitempty"`
	UpdatedAt Time            `json:"updated_at,omitempty"`
	DeletedAt Time            `json:"deleted_at,omitempty"`
	Order     string          `json:"order,omitempty"`
	Select    string          `json:"select,omitempty"`
	Offset    int             `json:"offset,omitempty"`
	Limit     int             `json:"limit,omitempty"`
	Deleted   bool            `json:"deleted,omitempty"`
}
