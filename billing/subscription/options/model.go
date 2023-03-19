package options

import (
	"github.com/driver005/gateway/core"
)

// SubscriptionsResourcePaymentMethodOptions
type SubscriptionsResourcePaymentMethodOptions struct {
	core.Model

	AcssDebit       *SubscriptionsResourcePaymentMethodOptionsAcssDebit       `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact      *SubscriptionsResourcePaymentMethodOptionsBancontact      `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card            *SubscriptionsResourcePaymentMethodOptionsCard            `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance *SubscriptionsResourcePaymentMethodOptionsCustomerBalance `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini         *SubscriptionsResourcePaymentMethodOptionsKonbini         `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount   *SubscriptionsResourcePaymentMethodOptionsUsBankAccount   `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
