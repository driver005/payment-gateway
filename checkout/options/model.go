package options

import "github.com/driver005/gateway/core"

// CheckoutSessionPaymentMethodOptions
type CheckoutSessionPaymentMethodOptions struct {
	core.Model

	AcssDebit        *CheckoutAcssDebitPaymentMethodOptions        `json:"acss_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Affirm           *CheckoutAffirmPaymentMethodOptions           `json:"affirm,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AfterpayClearpay *CheckoutAfterpayClearpayPaymentMethodOptions `json:"afterpay_clearpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Alipay           *CheckoutAlipayPaymentMethodOptions           `json:"alipay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	AuBecsDebit      *CheckoutAuBecsDebitPaymentMethodOptions      `json:"au_becs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	BacsDebit        *CheckoutBacsDebitPaymentMethodOptions        `json:"bacs_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Bancontact       *CheckoutBancontactPaymentMethodOptions       `json:"bancontact,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Boleto           *CheckoutBoletoPaymentMethodOptions           `json:"boleto,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Card             *CheckoutCardPaymentMethodOptions             `json:"card,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	CustomerBalance  *CheckoutCustomerBalancePaymentMethodOptions  `json:"customer_balance,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Eps              *CheckoutEpsPaymentMethodOptions              `json:"eps,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Fpx              *CheckoutFpxPaymentMethodOptions              `json:"fpx,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Giropay          *CheckoutGiropayPaymentMethodOptions          `json:"giropay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Grabpay          *CheckoutGrabPayPaymentMethodOptions          `json:"grabpay,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Ideal            *CheckoutIdealPaymentMethodOptions            `json:"ideal,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Klarna           *CheckoutKlarnaPaymentMethodOptions           `json:"klarna,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Konbini          *CheckoutKonbiniPaymentMethodOptions          `json:"konbini,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Oxxo             *CheckoutOxxoPaymentMethodOptions             `json:"oxxo,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	P24              *CheckoutP24PaymentMethodOptions              `json:"p24,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Paynow           *CheckoutPaynowPaymentMethodOptions           `json:"paynow,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Pix              *CheckoutPixPaymentMethodOptions              `json:"pix,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SepaDebit        *CheckoutSepaDebitPaymentMethodOptions        `json:"sepa_debit,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Sofort           *CheckoutSofortPaymentMethodOptions           `json:"sofort,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	UsBankAccount    *CheckoutUsBankAccountPaymentMethodOptions    `json:"us_bank_account,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}
