package method

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/radar"
)

type PaymentMethod struct {
	core.Model

	AcssDebit        *methods.PaymentMethodAcssDebit        `json:"acss_debit,omitempty"`
	Affirm           *methods.PaymentMethodAffirm           `json:"affirm,omitempty"`
	AfterpayClearpay *methods.PaymentMethodAfterpayClearpay `json:"afterpay_clearpay,omitempty"`
	Alipay           *methods.PaymentMethodAlipay           `json:"alipay,omitempty"`
	AuBecsDebit      *methods.PaymentMethodAuBecsDebit      `json:"au_becs_debit,omitempty"`
	BacsDebit        *methods.PaymentMethodBacsDebit        `json:"bacs_debit,omitempty"`
	Bancontact       *methods.PaymentMethodBancontact       `json:"bancontact,omitempty"`
	BillingDetails   *address.BillingDetails                `json:"billing_details"`
	Blik             *methods.PaymentMethodBlik             `json:"blik,omitempty"`
	Boleto           *methods.PaymentMethodBoleto           `json:"boleto,omitempty"`
	Card             *methods.PaymentMethodCard             `json:"card,omitempty"`
	CardPresent      *methods.PaymentMethodCardPresent      `json:"card_present,omitempty"`
	Customer         *methods.PaymentMethodCustomer         `json:"customer,omitempty"`
	CustomerBalance  *methods.PaymentMethodCustomerBalance  `json:"customer_balance,omitempty"`
	Eps              *methods.PaymentMethodEps              `json:"eps,omitempty"`
	Fpx              *methods.PaymentMethodFpx              `json:"fpx,omitempty"`
	Giropay          *methods.PaymentMethodGiropay          `json:"giropay,omitempty"`
	Grabpay          *methods.PaymentMethodGrabpay          `json:"grabpay,omitempty"`
	Ideal            *methods.PaymentMethodIdeal            `json:"ideal,omitempty"`
	InteracPresent   *methods.PaymentMethodInteracPresent   `json:"interac_present,omitempty"`
	Klarna           *methods.PaymentMethodKlarna           `json:"klarna,omitempty"`
	Konbini          *methods.PaymentMethodKonbini          `json:"konbini,omitempty"`
	Link             *methods.PaymentMethodLink             `json:"link,omitempty"`
	Oxxo             *methods.PaymentMethodOxxo             `json:"oxxo,omitempty"`
	P24              *methods.PaymentMethodP24              `json:"p24,omitempty"`
	Paynow           *methods.PaymentMethodPaynow           `json:"paynow,omitempty"`
	Pix              *methods.PaymentMethodPix              `json:"pix,omitempty"`
	Promptpay        *methods.PaymentMethodPromptpay        `json:"promptpay,omitempty"`
	RadarOptions     *radar.RadarOptions                    `json:"radar_options,omitempty"`
	SepaDebit        *methods.PaymentMethodSepaDebit        `json:"sepa_debit,omitempty"`
	Sofort           *methods.PaymentMethodSofort           `json:"sofort,omitempty"`
	UsBankAccount    *methods.PaymentMethodUsBankAccount    `json:"us_bank_account,omitempty"`
	WechatPay        *methods.PaymentMethodWechatPay        `json:"wechat_pay,omitempty"`
}
