package account

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&Account{},
	)
	if err != nil {
		panic(err)
	}
}