package tax

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&AutomaticTax{},
		&TaxDeductedAtSource{},
		&TaxCode{},
		&TaxRate{},
		&TaxId{},
	)
	if err != nil {
		panic(err)
	}
}
