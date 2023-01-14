package service

import (
	"context"

	"github.com/driver005/database"
	"github.com/driver005/gateway/logger"
	"github.com/rs/xid"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	Logger() *logger.Logger
}

func NewHandler() *Handler {
	return &Handler{}
}

func Paginate(offset int, pageSize int) func(db *database.DB) *database.DB {
	return func(db *database.DB) *database.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *Handler) newID(descriptor string) string {
	return descriptor + "_" + xid.New().String()
}
