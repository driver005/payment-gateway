package options

import "github.com/driver005/gateway/core"

// CheckoutSessionPaymentMethodOptions
type CheckoutSessionPaymentMethodOptions struct {
	core.Model

	AcssDebit        *CheckoutAcssDebitPaymentMethodOptions        `json:"acss_debit,omitempty" database:"foreignKey:id"`
	Affirm           *CheckoutAffirmPaymentMethodOptions           `json:"affirm,omitempty" database:"foreignKey:id"`
	AfterpayClearpay *CheckoutAfterpayClearpayPaymentMethodOptions `json:"afterpay_clearpay,omitempty" database:"foreignKey:id"`
	Alipay           *CheckoutAlipayPaymentMethodOptions           `json:"alipay,omitempty" database:"foreignKey:id"`
	AuBecsDebit      *CheckoutAuBecsDebitPaymentMethodOptions      `json:"au_becs_debit,omitempty" database:"foreignKey:id"`
	BacsDebit        *CheckoutBacsDebitPaymentMethodOptions        `json:"bacs_debit,omitempty" database:"foreignKey:id"`
	Bancontact       *CheckoutBancontactPaymentMethodOptions       `json:"bancontact,omitempty" database:"foreignKey:id"`
	Boleto           *CheckoutBoletoPaymentMethodOptions           `json:"boleto,omitempty" database:"foreignKey:id"`
	Card             *CheckoutCardPaymentMethodOptions             `json:"card,omitempty" database:"foreignKey:id"`
	CustomerBalance  *CheckoutCustomerBalancePaymentMethodOptions  `json:"customer_balance,omitempty" database:"foreignKey:id"`
	Eps              *CheckoutEpsPaymentMethodOptions              `json:"eps,omitempty" database:"foreignKey:id"`
	Fpx              *CheckoutFpxPaymentMethodOptions              `json:"fpx,omitempty" database:"foreignKey:id"`
	Giropay          *CheckoutGiropayPaymentMethodOptions          `json:"giropay,omitempty" database:"foreignKey:id"`
	Grabpay          *CheckoutGrabPayPaymentMethodOptions          `json:"grabpay,omitempty" database:"foreignKey:id"`
	Ideal            *CheckoutIdealPaymentMethodOptions            `json:"ideal,omitempty" database:"foreignKey:id"`
	Klarna           *CheckoutKlarnaPaymentMethodOptions           `json:"klarna,omitempty" database:"foreignKey:id"`
	Konbini          *CheckoutKonbiniPaymentMethodOptions          `json:"konbini,omitempty" database:"foreignKey:id"`
	Oxxo             *CheckoutOxxoPaymentMethodOptions             `json:"oxxo,omitempty" database:"foreignKey:id"`
	P24              *CheckoutP24PaymentMethodOptions              `json:"p24,omitempty" database:"foreignKey:id"`
	Paynow           *CheckoutPaynowPaymentMethodOptions           `json:"paynow,omitempty" database:"foreignKey:id"`
	Pix              *CheckoutPixPaymentMethodOptions              `json:"pix,omitempty" database:"foreignKey:id"`
	SepaDebit        *CheckoutSepaDebitPaymentMethodOptions        `json:"sepa_debit,omitempty" database:"foreignKey:id"`
	Sofort           *CheckoutSofortPaymentMethodOptions           `json:"sofort,omitempty" database:"foreignKey:id"`
	UsBankAccount    *CheckoutUsBankAccountPaymentMethodOptions    `json:"us_bank_account,omitempty" database:"foreignKey:id"`
}
