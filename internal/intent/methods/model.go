package methods

import (
	"github.com/driver005/gateway/core"
)

// PaymentIntentPaymentMethodOptions
type PaymentIntentPaymentMethodOptions struct {
	core.Model

	AcssDebit        *PaymentIntentPaymentMethodOptionsAcssDebit1       `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Affirm           *PaymentIntentPaymentMethodOptionsAffirm           `json:"affirm,omitempty" database:"foreignKey:id"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay,omitempty" database:"foreignKey:id"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipay           `json:"alipay,omitempty" database:"foreignKey:id"`
	AuBecsDebit      *PaymentIntentPaymentMethodOptionsAuBecsDebit1     `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	BacsDebit        *PaymentIntentPaymentMethodOptionsBacsDebit        `json:"bacs_debit,omitempty" database:"foreignKey:id"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontact       `json:"bancontact,omitempty" database:"foreignKey:id"`
	Blik             *PaymentIntentPaymentMethodOptionsBlik             `json:"blik,omitempty" database:"foreignKey:id"`
	Boleto           *PaymentIntentPaymentMethodOptionsBoleto           `json:"boleto,omitempty" database:"foreignKey:id"`
	Card             *PaymentIntentPaymentMethodOptionsCard1            `json:"card,omitempty" database:"foreignKey:id"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresent      `json:"card_present,omitempty" database:"foreignKey:id"`
	CustomerBalance  *PaymentIntentPaymentMethodOptionsCustomerBalance  `json:"customer_balance,omitempty" database:"foreignKey:id"`
	Eps              *PaymentIntentPaymentMethodOptionsEps1             `json:"eps,omitempty" database:"foreignKey:id"`
	Fpx              *PaymentIntentPaymentMethodOptionsFpx              `json:"fpx,omitempty" database:"foreignKey:id"`
	Giropay          *PaymentIntentPaymentMethodOptionsGiropay          `json:"giropay,omitempty" database:"foreignKey:id"`
	Grabpay          *PaymentIntentPaymentMethodOptionsGrabpay          `json:"grabpay,omitempty" database:"foreignKey:id"`
	Ideal            *PaymentIntentPaymentMethodOptionsIdeal            `json:"ideal,omitempty" database:"foreignKey:id"`
	InteracPresent   *PaymentIntentPaymentMethodOptionsInteracPresent   `json:"interac_present,omitempty" database:"foreignKey:id"`
	Klarna           *PaymentIntentPaymentMethodOptionsKlarna           `json:"klarna,omitempty" database:"foreignKey:id"`
	Konbini          *PaymentIntentPaymentMethodOptionsKonbini          `json:"konbini,omitempty" database:"foreignKey:id"`
	Link             *PaymentIntentPaymentMethodOptionsLink1            `json:"link,omitempty" database:"foreignKey:id"`
	Oxxo             *PaymentIntentPaymentMethodOptionsOxxo             `json:"oxxo,omitempty" database:"foreignKey:id"`
	BtcPay           *PaymentIntentPaymentMethodOptionsBtcPay           `json:"btc_pay,omitempty" database:"foreignKey:id"`
	P24              *PaymentIntentPaymentMethodOptionsP24              `json:"p24,omitempty" database:"foreignKey:id"`
	Paynow           *PaymentIntentPaymentMethodOptionsPaynow           `json:"paynow,omitempty" database:"foreignKey:id"`
	Pix              *PaymentIntentPaymentMethodOptionsPix              `json:"pix,omitempty" database:"foreignKey:id"`
	Promptpay        *PaymentIntentPaymentMethodOptionsPromptpay        `json:"promptpay,omitempty" database:"foreignKey:id"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebit1       `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofort           `json:"sofort,omitempty" database:"foreignKey:id"`
	UsBankAccount    *PaymentIntentPaymentMethodOptionsUsBankAccount1   `json:"us_bank_account,omitempty" database:"foreignKey:id"`
	WechatPay        *PaymentIntentPaymentMethodOptionsWechatPay        `json:"wechat_pay,omitempty" database:"foreignKey:id"`
}
