package details

import "github.com/driver005/gateway/core"

type PaymentMethodDetails struct {
	core.Model

	AchCreditTransfer *PaymentMethodDetailsAchCreditTransfer `json:"ach_credit_transfer,omitempty"`
	AchDebit          *PaymentMethodDetailsAchDebit          `json:"ach_debit,omitempty"`
	AcssDebit         *PaymentMethodDetailsAcssDebit         `json:"acss_debit,omitempty"`
	Affirm            *PaymentMethodDetailsAffirm            `json:"affirm,omitempty"`
	AfterpayClearpay  *PaymentMethodDetailsAfterpayClearpay  `json:"afterpay_clearpay,omitempty"`
	Alipay            *PaymentMethodDetailsAlipay            `json:"alipay,omitempty"`
	AuBecsDebit       *PaymentMethodDetailsAuBecsDebit       `json:"au_becs_debit,omitempty"`
	BacsDebit         *PaymentMethodDetailsBacsDebit         `json:"bacs_debit,omitempty"`
	Bancontact        *PaymentMethodDetailsBancontact        `json:"bancontact,omitempty"`
	Blik              *PaymentMethodDetailsBlik              `json:"blik,omitempty"`
	Boleto            *PaymentMethodDetailsBoleto            `json:"boleto,omitempty"`
	Card              *PaymentMethodDetailsCard              `json:"card,omitempty"`
	CardPresent       *PaymentMethodDetailsCardPresent       `json:"card_present,omitempty"`
	CustomerBalance   *PaymentMethodDetailsCustomerBalance   `json:"customer_balance,omitempty"`
	Eps               *PaymentMethodDetailsEps               `json:"eps,omitempty"`
	Fpx               *PaymentMethodDetailsFpx               `json:"fpx,omitempty"`
	Giropay           *PaymentMethodDetailsGiropay           `json:"giropay,omitempty"`
	Grabpay           *PaymentMethodDetailsGrabpay           `json:"grabpay,omitempty"`
	Ideal             *PaymentMethodDetailsIdeal             `json:"ideal,omitempty"`
	InteracPresent    *PaymentMethodDetailsInteracPresent    `json:"interac_present,omitempty"`
	Klarna            *PaymentMethodDetailsKlarna            `json:"klarna,omitempty"`
	Konbini           *PaymentMethodDetailsKonbini           `json:"konbini,omitempty"`
	Link              *PaymentMethodDetailsLink              `json:"link,omitempty"`
	Multibanco        *PaymentMethodDetailsMultibanco        `json:"multibanco,omitempty"`
	Oxxo              *PaymentMethodDetailsOxxo              `json:"oxxo,omitempty"`
	P24               *PaymentMethodDetailsP24               `json:"p24,omitempty"`
	Paynow            *PaymentMethodDetailsPaynow            `json:"paynow,omitempty"`
	Pix               *PaymentMethodDetailsPix               `json:"pix,omitempty"`
	Promptpay         *PaymentMethodDetailsPromptpay         `json:"promptpay,omitempty"`
	SepaDebit         *PaymentMethodDetailsSepaDebit         `json:"sepa_debit,omitempty"`
	Sofort            *PaymentMethodDetailsSofort            `json:"sofort,omitempty"`
	StripeAccount     *PaymentMethodDetailsStripeAccount     `json:"stripe_account,omitempty"`
	UsBankAccount     *PaymentMethodDetailsUsBankAccount     `json:"us_bank_account,omitempty"`
	Wechat            *PaymentMethodDetailsWechat            `json:"wechat,omitempty"`
	WechatPay         *PaymentMethodDetailsWechatPay         `json:"wechat_pay,omitempty"`
}
