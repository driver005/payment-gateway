package database

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
		db   *database.DB
		l    *logger.Logger
	}
	// Dependencies interface {
	// 	Tracer(ctx context.Context) trace.Tracer
	// }
	transactionContextType string
)

func NewManager(db *database.DB, l *logger.Logger) (*Manager, error) {

	return &Manager{
		db:   db,
		l:    l,
	}, nil
}

// func (p *Manager) BeginTX(ctx context.Context) (context.Context, error) {
// 	_, ok := ctx.Value(transactionContextKey).(*pop.Connection)
// 	if ok {
// 		return ctx, helper.WithStack(ErrTransactionOpen)
// 	}

// 	tx, err := p.conn.Store.TransactionContextOptions(ctx, &sql.TxOptions{
// 		Isolation: sql.LevelRepeatableRead,
// 		ReadOnly:  false,
// 	})
// 	c := &pop.Connection{
// 		TX:      tx,
// 		Store:   tx,
// 		ID:      helper.String(30),
// 		Dialect: p.conn.Dialect,
// 	}
// 	return context.WithValue(ctx, transactionContextKey, c), err
// }

// func (p *Manager) Commit(ctx context.Context) error {
// 	c, ok := ctx.Value(transactionContextKey).(*pop.Connection)
// 	if !ok || c.TX == nil {
// 		return helper.WithStack(ErrNoTransactionOpen)
// 	}

// 	return c.TX.Commit()
// }

// func (p *Manager) Rollback(ctx context.Context) error {
// 	c, ok := ctx.Value(transactionContextKey).(*pop.Connection)
// 	if !ok || c.TX == nil {
// 		return helper.WithStack(ErrNoTransactionOpen)
// 	}

// 	return c.TX.Rollback()
// }

func (p *Manager) DB(ctx context.Context) *database.DB {
	if c, ok := ctx.Value(transactionContextKey).(*database.DB); ok {
		return c.WithContext(ctx)
	}
	return p.db.WithContext(ctx)
}

// func (p *Manager) Transaction(ctx context.Context, f func(ctx context.Context, c *pop.Connection) error) error {
// 	isNested := true
// 	c, ok := ctx.Value(transactionContextKey).(*pop.Connection)
// 	if !ok {
// 		isNested = false

// 		var err error
// 		c, err = p.conn.WithContext(ctx).NewTransaction()

// 		if err != nil {
// 			return helper.WithStack(err)
// 		}
// 	}

// 	if err := f(context.WithValue(ctx, transactionContextKey, c), c); err != nil {
// 		if !isNested {
// 			if err := c.TX.Rollback(); err != nil {
// 				return helper.WithStack(err)
// 			}
// 		}
// 		return err
// 	}

// 	// commit if there is no wrapping transaction
// 	if !isNested {
// 		return helper.WithStack(c.TX.Commit())
// 	}

// 	return nil
// }
