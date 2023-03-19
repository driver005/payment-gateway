package methods

import (
	"github.com/driver005/gateway/core"
)

type SetupAttemptPaymentMethodDetails struct {
	core.Model

	AcssDebit     core.JSONB                                   `json:"acss_debit,omitempty"`
	AuBecsDebit   core.JSONB                                   `json:"au_becs_debit,omitempty"`
	BacsDebit     core.JSONB                                   `json:"bacs_debit,omitempty"`
	Bancontact    *SetupAttemptPaymentMethodDetailsBancontact  `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Blik          core.JSONB                                   `json:"blik,omitempty"`
	Boleto        core.JSONB                                   `json:"boleto,omitempty"`
	Card          *SetupAttemptPaymentMethodDetailsCard        `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CardPresent   *SetupAttemptPaymentMethodDetailsCardPresent `json:"card_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Ideal         *SetupAttemptPaymentMethodDetailsIdeal       `json:"ideal,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Klarna        core.JSONB                                   `json:"klarna,omitempty"`
	Link          core.JSONB                                   `json:"link,omitempty"`
	SepaDebit     core.JSONB                                   `json:"sepa_debit,omitempty"`
	Sofort        *SetupAttemptPaymentMethodDetailsSofort      `json:"sofort,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Type          string                                       `json:"type,omitempty"`
	UsBankAccount core.JSONB                                   `json:"us_bank_account,omitempty"`
}
