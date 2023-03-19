package details

import (
	"github.com/driver005/gateway/core"
)

type PaymentMethodDetails struct {
	core.Model

	AchCreditTransfer *PaymentMethodDetailsAchCreditTransfer `json:"ach_credit_transfer,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AchDebit          *PaymentMethodDetailsAchDebit          `json:"ach_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AcssDebit         *PaymentMethodDetailsAcssDebit         `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Affirm            *PaymentMethodDetailsAffirm            `json:"affirm,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AfterpayClearpay  *PaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Alipay            *PaymentMethodDetailsAlipay            `json:"alipay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AuBecsDebit       *PaymentMethodDetailsAuBecsDebit       `json:"au_becs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BacsDebit         *PaymentMethodDetailsBacsDebit         `json:"bacs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact        *PaymentMethodDetailsBancontact        `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Blik              *PaymentMethodDetailsBlik              `json:"blik,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Boleto            *PaymentMethodDetailsBoleto            `json:"boleto,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card              *PaymentMethodDetailsCard              `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CardPresent       *PaymentMethodDetailsCardPresent       `json:"card_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance   *PaymentMethodDetailsCustomerBalance   `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Eps               *PaymentMethodDetailsEps               `json:"eps,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Fpx               *PaymentMethodDetailsFpx               `json:"fpx,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Giropay           *PaymentMethodDetailsGiropay           `json:"giropay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Grabpay           *PaymentMethodDetailsGrabpay           `json:"grabpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Ideal             *PaymentMethodDetailsIdeal             `json:"ideal,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	InteracPresent    *PaymentMethodDetailsInteracPresent    `json:"interac_present,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Klarna            *PaymentMethodDetailsKlarna            `json:"klarna,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini           *PaymentMethodDetailsKonbini           `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Link              *PaymentMethodDetailsLink              `json:"link,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Multibanco        *PaymentMethodDetailsMultibanco        `json:"multibanco,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Oxxo              *PaymentMethodDetailsOxxo              `json:"oxxo,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	P24               *PaymentMethodDetailsP24               `json:"p24,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Paynow            *PaymentMethodDetailsPaynow            `json:"paynow,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Pix               *PaymentMethodDetailsPix               `json:"pix,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Promptpay         *PaymentMethodDetailsPromptpay         `json:"promptpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SepaDebit         *PaymentMethodDetailsSepaDebit         `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Sofort            *PaymentMethodDetailsSofort            `json:"sofort,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	StripeAccount     *PaymentMethodDetailsStripeAccount     `json:"stripe_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount     *PaymentMethodDetailsUsBankAccount     `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Wechat            *PaymentMethodDetailsWechat            `json:"wechat,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	WechatPay         *PaymentMethodDetailsWechatPay         `json:"wechat_pay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
