package core

import (
	"github.com/gofrs/uuid"
)

type Filter struct {
	Id        []uuid.NullUUID `json:"id,omitempty"`
	CreatedAt Time            `json:"created_at,omitempty"`
	UpdatedAt Time            `json:"updated_at,omitempty"`
}
