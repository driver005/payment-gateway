package method

import (
	"context"

	"github.com/driver005/database"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/logger"
	"github.com/driver005/gateway/service"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	ClientManager() *db.Handler
	Service() *service.Handler
	Logger() *logger.Logger
}

func NewHandler(r Registry, host string) *Handler {
	return &Handler{
		r: r,
	}
}
