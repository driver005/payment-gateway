package database

import (
	"context"

	"github.com/driver005/database"
)

type (
	Database interface {
		DB(ctx context.Context) *database.DB
	}
	Provider interface {
		Database() Database
	}
)
