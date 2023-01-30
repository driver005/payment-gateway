package details

import (
	"github.com/driver005/gateway/core"
)

type PaymentMethodDetails struct {
	core.Model

	AchCreditTransfer *PaymentMethodDetailsAchCreditTransfer `json:"ach_credit_transfer,omitempty" database:"foreignKey:id"`
	AchDebit          *PaymentMethodDetailsAchDebit          `json:"ach_debit,omitempty" database:"foreignKey:id"`
	AcssDebit         *PaymentMethodDetailsAcssDebit         `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Affirm            *PaymentMethodDetailsAffirm            `json:"affirm,omitempty" database:"foreignKey:id"`
	AfterpayClearpay  *PaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay,omitempty" database:"foreignKey:id"`
	Alipay            *PaymentMethodDetailsAlipay            `json:"alipay,omitempty" database:"foreignKey:id"`
	AuBecsDebit       *PaymentMethodDetailsAuBecsDebit       `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	BacsDebit         *PaymentMethodDetailsBacsDebit         `json:"bacs_debit,omitempty" database:"foreignKey:id"`
	Bancontact        *PaymentMethodDetailsBancontact        `json:"bancontact,omitempty" database:"foreignKey:id"`
	Blik              *PaymentMethodDetailsBlik              `json:"blik,omitempty" database:"foreignKey:id"`
	Boleto            *PaymentMethodDetailsBoleto            `json:"boleto,omitempty" database:"foreignKey:id"`
	Card              *PaymentMethodDetailsCard              `json:"card,omitempty" database:"foreignKey:id"`
	CardPresent       *PaymentMethodDetailsCardPresent       `json:"card_present,omitempty" database:"foreignKey:id"`
	CustomerBalance   *PaymentMethodDetailsCustomerBalance   `json:"customer_balance,omitempty" database:"foreignKey:id"`
	Eps               *PaymentMethodDetailsEps               `json:"eps,omitempty" database:"foreignKey:id"`
	Fpx               *PaymentMethodDetailsFpx               `json:"fpx,omitempty" database:"foreignKey:id"`
	Giropay           *PaymentMethodDetailsGiropay           `json:"giropay,omitempty" database:"foreignKey:id"`
	Grabpay           *PaymentMethodDetailsGrabpay           `json:"grabpay,omitempty" database:"foreignKey:id"`
	Ideal             *PaymentMethodDetailsIdeal             `json:"ideal,omitempty" database:"foreignKey:id"`
	InteracPresent    *PaymentMethodDetailsInteracPresent    `json:"interac_present,omitempty" database:"foreignKey:id"`
	Klarna            *PaymentMethodDetailsKlarna            `json:"klarna,omitempty" database:"foreignKey:id"`
	Konbini           *PaymentMethodDetailsKonbini           `json:"konbini,omitempty" database:"foreignKey:id"`
	Link              *PaymentMethodDetailsLink              `json:"link,omitempty" database:"foreignKey:id"`
	Multibanco        *PaymentMethodDetailsMultibanco        `json:"multibanco,omitempty" database:"foreignKey:id"`
	Oxxo              *PaymentMethodDetailsOxxo              `json:"oxxo,omitempty" database:"foreignKey:id"`
	P24               *PaymentMethodDetailsP24               `json:"p24,omitempty" database:"foreignKey:id"`
	Paynow            *PaymentMethodDetailsPaynow            `json:"paynow,omitempty" database:"foreignKey:id"`
	Pix               *PaymentMethodDetailsPix               `json:"pix,omitempty" database:"foreignKey:id"`
	Promptpay         *PaymentMethodDetailsPromptpay         `json:"promptpay,omitempty" database:"foreignKey:id"`
	SepaDebit         *PaymentMethodDetailsSepaDebit         `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	Sofort            *PaymentMethodDetailsSofort            `json:"sofort,omitempty" database:"foreignKey:id"`
	StripeAccount     *PaymentMethodDetailsStripeAccount     `json:"stripe_account,omitempty" database:"foreignKey:id"`
	UsBankAccount     *PaymentMethodDetailsUsBankAccount     `json:"us_bank_account,omitempty" database:"foreignKey:id"`
	Wechat            *PaymentMethodDetailsWechat            `json:"wechat,omitempty" database:"foreignKey:id"`
	WechatPay         *PaymentMethodDetailsWechatPay         `json:"wechat_pay,omitempty" database:"foreignKey:id"`
}
