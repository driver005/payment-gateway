package country

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(&Country{})
	if err != nil {
		panic(err)
	}
}
