package sql

import (
	"context"
	"embed"

	"github.com/driver005/database"
	"github.com/driver005/gateway/logger"

	"github.com/pkg/errors"
)

var Migrations embed.FS

const transactionContextKey transactionContextType = "transactionConnection"

var (
	ErrTransactionOpen   = errors.New("There is already a transaction in this context.")
	ErrNoTransactionOpen = errors.New("There is no transaction in this context.")
)

type (
	Manager struct {
		db *database.DB
		l  *logger.Logger
	}
	// Dependencies interface {
	// 	Tracer(ctx context.Context) trace.Tracer
	// }
	transactionContextType string
)

func NewManager(db *database.DB, l *logger.Logger) (*Manager, error) {
	return &Manager{
		db: db,
		l:  l,
	}, nil
}

func (p *Manager) DB(ctx context.Context) *database.DB {
	if c, ok := ctx.Value(transactionContextKey).(*database.DB); ok {
		return c.WithContext(ctx)
	}
	return p.db.WithContext(ctx)
}
