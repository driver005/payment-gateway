package rule

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(&Rule{})
	if err != nil {
		panic(err)
	}
}
