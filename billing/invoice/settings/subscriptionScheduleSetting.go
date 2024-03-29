package settings

import "github.com/driver005/gateway/core"

// InvoiceSettingSubscriptionScheduleSetting
type InvoiceSettingSubscriptionScheduleSetting struct {
	core.Model

	// Number of days within which a customer must pay invoices generated by this subscription schedule. This value will be `null` for subscription schedules where `billing=charge_automatically`.
	DaysUntilDue int `json:"days_until_due,omitempty"`
}
