package helper

import (
	"errors"

	"github.com/driver005/database"
	"github.com/lib/pq"
)

func ParseError(err error) error {
	pgError := err.(*pq.Error)

	if !errors.Is(err, pgError) {
		return WithStack(err)
	}

	return errors.New(pgError.Detail)
}

func Paginate(offset int, pageSize int) func(db *database.DB) *database.DB {
	return func(db *database.DB) *database.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}
