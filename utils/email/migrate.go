package email

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(&EmailSent{})
	if err != nil {
		panic(err)
	}
}