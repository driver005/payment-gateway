package methods

import (
	"github.com/driver005/gateway/core"
)

type MandatePaymentMethodDetails struct {
	core.Model

	AcssDebit     *MandateAcssDebit     `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AuBecsDebit   *MandateAuBecsDebit   `json:"au_becs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BacsDebit     *MandateBacsDebit     `json:"bacs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Blik          *MandateBlik          `json:"blik,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card          *MandateCard          `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Link          *MandateLink          `json:"link,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SepaDebit     *MandateSepaDebit     `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount *MandateUsBankAccount `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Type          string                `json:"type,omitempty"`
}
