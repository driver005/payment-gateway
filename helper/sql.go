package helper

import (
	"errors"

	"github.com/lib/pq"
)

func ParseError(err error) error {
	pgError := err.(*pq.Error)

	if !errors.Is(err, pgError) {
		return WithStack(err)
	}

	return errors.New(pgError.Detail)
}
