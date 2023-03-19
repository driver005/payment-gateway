package methods

import (
	"github.com/driver005/gateway/core"
)

// PaymentIntentPaymentMethodOptions
type PaymentIntentPaymentMethodOptions struct {
	core.Model

	AcssDebit        *PaymentIntentPaymentMethodOptionsAcssDebit1       `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Affirm           *PaymentIntentPaymentMethodOptionsAffirm           `json:"affirm,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipay           `json:"alipay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AuBecsDebit      *PaymentIntentPaymentMethodOptionsAuBecsDebit1     `json:"au_becs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BacsDebit        *PaymentIntentPaymentMethodOptionsBacsDebit        `json:"bacs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontact       `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Blik             *PaymentIntentPaymentMethodOptionsBlik             `json:"blik,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Boleto           *PaymentIntentPaymentMethodOptionsBoleto           `json:"boleto,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card             *PaymentIntentPaymentMethodOptionsCard1            `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresent      `json:"card_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance  *PaymentIntentPaymentMethodOptionsCustomerBalance  `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Eps              *PaymentIntentPaymentMethodOptionsEps1             `json:"eps,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Fpx              *PaymentIntentPaymentMethodOptionsFpx              `json:"fpx,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Giropay          *PaymentIntentPaymentMethodOptionsGiropay          `json:"giropay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Grabpay          *PaymentIntentPaymentMethodOptionsGrabpay          `json:"grabpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Ideal            *PaymentIntentPaymentMethodOptionsIdeal            `json:"ideal,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	InteracPresent   *PaymentIntentPaymentMethodOptionsInteracPresent   `json:"interac_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Klarna           *PaymentIntentPaymentMethodOptionsKlarna           `json:"klarna,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini          *PaymentIntentPaymentMethodOptionsKonbini          `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Link             *PaymentIntentPaymentMethodOptionsLink1            `json:"link,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Oxxo             *PaymentIntentPaymentMethodOptionsOxxo             `json:"oxxo,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BtcPay           *PaymentIntentPaymentMethodOptionsBtcPay           `json:"btc_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	P24              *PaymentIntentPaymentMethodOptionsP24              `json:"p24,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Paynow           *PaymentIntentPaymentMethodOptionsPaynow           `json:"paynow,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Pix              *PaymentIntentPaymentMethodOptionsPix              `json:"pix,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Promptpay        *PaymentIntentPaymentMethodOptionsPromptpay        `json:"promptpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebit1       `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofort           `json:"sofort,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount    *PaymentIntentPaymentMethodOptionsUsBankAccount1   `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPay        *PaymentIntentPaymentMethodOptionsWechatPay        `json:"wechat_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
