package core

import (
	"time"

	"github.com/driver005/database"
	"github.com/gofrs/uuid"
)

type Model struct {
	Id        uuid.UUID          `json:"id" database:"primarykey"`
	CreatedAt time.Time          `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt database.DeletedAt `json:"deleted_at,omitempty" db:"deleted_at"`
}

func (a *Model) BeforeCreate(tx *database.DB) (err error) {
	a.Id, err = uuid.NewV4()
	if err != nil {
		return err
	}
	return nil
}
