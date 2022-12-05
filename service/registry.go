package service

import (
	"context"

	"github.com/driver005/database"
	db "github.com/driver005/gateway/database"
	"github.com/driver005/gateway/repository"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
	ClientManager() *db.Handler
	Repository() repository.TransactionRepository
}

func NewHandler(r Registry) *Handler {
	return &Handler{r: r}
}
