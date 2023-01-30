package methods

import (
	"github.com/driver005/gateway/core"
)

type MandatePaymentMethodDetails struct {
	core.Model

	AcssDebit     *MandateAcssDebit     `json:"acss_debit,omitempty" database:"foreignKey:id"`
	AuBecsDebit   *MandateAuBecsDebit   `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	BacsDebit     *MandateBacsDebit     `json:"bacs_debit,omitempty" database:"foreignKey:id"`
	Blik          *MandateBlik          `json:"blik,omitempty" database:"foreignKey:id"`
	Card          *MandateCard          `json:"card,omitempty" database:"foreignKey:id"`
	Link          *MandateLink          `json:"link,omitempty" database:"foreignKey:id"`
	SepaDebit     *MandateSepaDebit     `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	UsBankAccount *MandateUsBankAccount `json:"us_bank_account,omitempty" database:"foreignKey:id"`
	Type          string                `json:"type,omitempty"`
}
