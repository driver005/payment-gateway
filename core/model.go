package core

import (
	"time"

	"github.com/driver005/database"
	"github.com/google/uuid"
)

type Model struct {
	Id        uuid.UUID          `json:"id" database:"primary_key"`
	Object    string             `json:"object"`
	Livemode  bool               `json:"livemode"`
	Metadata  JSONB              `json:"metadata,omitempty" database:"default:null"`
	CreatedAt time.Time          `json:"created_at,omitempty" db:"created_at"`
	UpdatedAt time.Time          `json:"updated_at,omitempty" db:"updated_at"`
	DeletedAt database.DeletedAt `json:"deleted_at,omitempty" db:"deleted_at" swaggerignore:"true"`
}

func (m *Model) BeforeCreate(tx *database.DB) (err error) {
	if m.Id == uuid.Nil {
		m.Id, err = uuid.NewUUID()
		if err != nil {
			return err
		}
	}

	if m.Object == "" {
		m.Object = tx.Statement.Schema.Table
	}

	m.CreatedAt = time.Now().UTC().Round(time.Second)
	m.UpdatedAt = m.CreatedAt

	return nil
}

func (m *Model) BeforeUpdate(tx *database.DB) (err error) {
	m.UpdatedAt = time.Now().UTC().Round(time.Second)

	return nil
}
