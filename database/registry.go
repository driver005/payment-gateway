package database

import (
	"context"

	"github.com/driver005/database"
)

type Handler struct {
	r Registry
}

func (h *Handler) remove() {

}

type Registry interface {
	Manager(ctx context.Context) *database.DB
}

func Paginate(offset int, pageSize int) func(db *database.DB) *database.DB {
	return func(db *database.DB) *database.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func NewHandler(r Registry) *Handler {
	return &Handler{r: r}
}
