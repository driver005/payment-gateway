package errors

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&ApiErrors{},
	)
	if err != nil {
		panic(err)
	}
}
