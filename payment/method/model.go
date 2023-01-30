package method

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/internal/customer"
	"github.com/driver005/gateway/payment/method/methods"
	"github.com/driver005/gateway/utils/address"
	"github.com/driver005/gateway/utils/radar"
	"github.com/google/uuid"
)

type PaymentMethod struct {
	core.Model

	AcssDebit        *methods.PaymentMethodAcssDebit        `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Affirm           *methods.PaymentMethodAffirm           `json:"affirm,omitempty" database:"foreignKey:id"`
	AfterpayClearpay *methods.PaymentMethodAfterpayClearpay `json:"afterpay_clearpay,omitempty" database:"foreignKey:id"`
	Alipay           *methods.PaymentMethodAlipay           `json:"alipay,omitempty" database:"foreignKey:id"`
	AuBecsDebit      *methods.PaymentMethodAuBecsDebit      `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	BacsDebit        *methods.PaymentMethodBacsDebit        `json:"bacs_debit,omitempty" database:"foreignKey:id"`
	Bancontact       *methods.PaymentMethodBancontact       `json:"bancontact,omitempty" database:"foreignKey:id"`
	BillingDetails   *address.BillingDetails                `json:"billing_details,omitempty" database:"foreignKey:id"`
	Blik             *methods.PaymentMethodBlik             `json:"blik,omitempty" database:"foreignKey:id"`
	Boleto           *methods.PaymentMethodBoleto           `json:"boleto,omitempty" database:"foreignKey:id"`
	BtcPay           *methods.PaymentMethodBtcPay           `json:"btc_pay,omitempty" database:"foreignKey:id"`
	Card             *methods.PaymentMethodCard             `json:"card,omitempty" database:"foreignKey:id"`
	CardPresent      *methods.PaymentMethodCardPresent      `json:"card_present,omitempty" database:"foreignKey:id"`
	CustomerBalance  *methods.PaymentMethodCustomerBalance  `json:"customer_balance,omitempty" database:"foreignKey:id"`
	Eps              *methods.PaymentMethodEps              `json:"eps,omitempty" database:"foreignKey:id"`
	Fpx              *methods.PaymentMethodFpx              `json:"fpx,omitempty" database:"foreignKey:id"`
	Giropay          *methods.PaymentMethodGiropay          `json:"giropay,omitempty" database:"foreignKey:id"`
	Grabpay          *methods.PaymentMethodGrabpay          `json:"grabpay,omitempty" database:"foreignKey:id"`
	Ideal            *methods.PaymentMethodIdeal            `json:"ideal,omitempty" database:"foreignKey:id"`
	InteracPresent   *methods.PaymentMethodInteracPresent   `json:"interac_present,omitempty" database:"foreignKey:id"`
	Klarna           *methods.PaymentMethodKlarna           `json:"klarna,omitempty" database:"foreignKey:id"`
	Konbini          *methods.PaymentMethodKonbini          `json:"konbini,omitempty" database:"foreignKey:id"`
	Link             *methods.PaymentMethodLink             `json:"link,omitempty" database:"foreignKey:id"`
	Oxxo             *methods.PaymentMethodOxxo             `json:"oxxo,omitempty" database:"foreignKey:id"`
	P24              *methods.PaymentMethodP24              `json:"p24,omitempty" database:"foreignKey:id"`
	Paynow           *methods.PaymentMethodPaynow           `json:"paynow,omitempty" database:"foreignKey:id"`
	Pix              *methods.PaymentMethodPix              `json:"pix,omitempty" database:"foreignKey:id"`
	Promptpay        *methods.PaymentMethodPromptpay        `json:"promptpay,omitempty" database:"foreignKey:id"`
	RadarOptions     *radar.RadarOptions                    `json:"radar_options,omitempty" database:"foreignKey:id"`
	SepaDebit        *methods.PaymentMethodSepaDebit        `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	Sofort           *methods.PaymentMethodSofort           `json:"sofort,omitempty" database:"foreignKey:id"`
	UsBankAccount    *methods.PaymentMethodUsBankAccount    `json:"us_bank_account,omitempty" database:"foreignKey:id"`
	WechatPay        *methods.PaymentMethodWechatPay        `json:"wechat_pay,omitempty" database:"foreignKey:id"`
	Type             string                                 `json:"type"`

	CustomerId uuid.UUID          `json:"customer_id"`
	Customer   *customer.Customer `json:"customer,omitempty" database:"foreignKey:customer_id"`
}
