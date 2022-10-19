package handler

import (
	"context"

	"github.com/driver005/database"
)

type Handler struct {
	r Registry
}

type Registry interface {
	Manager(ctx context.Context) *database.DB
}

func NewHandler(r Registry) *Handler {
	return &Handler{
		r: r,
	}
}
