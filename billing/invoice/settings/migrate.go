package settings

import "github.com/driver005/database"

type Registry interface {
	Context() *database.DB
}

func Migrate(r Registry) {
	err := r.Context().AutoMigrate(
		&InvoiceSettingCustomField{},
		&InvoiceSettingCustomerSetting{},
		&InvoiceSettingQuoteSetting{},
		&InvoiceSettingRenderingOptions{},
		&InvoiceSettingSubscriptionScheduleSetting{},
	)
	if err != nil {
		panic(err)
	}
}
