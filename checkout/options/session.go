package checkout

// CheckoutSessionPaymentMethodOptions
type CheckoutSessionPaymentMethodOptions struct {
	AcssDebit        *CheckoutAcssDebitPaymentMethodOptions        `json:"acss_debit,omitempty"`
	Affirm           *CheckoutAffirmPaymentMethodOptions           `json:"affirm,omitempty"`
	AfterpayClearpay *CheckoutAfterpayClearpayPaymentMethodOptions `json:"afterpay_clearpay,omitempty"`
	Alipay           *CheckoutAlipayPaymentMethodOptions           `json:"alipay,omitempty"`
	AuBecsDebit      *CheckoutAuBecsDebitPaymentMethodOptions      `json:"au_becs_debit,omitempty"`
	BacsDebit        *CheckoutBacsDebitPaymentMethodOptions        `json:"bacs_debit,omitempty"`
	Bancontact       *CheckoutBancontactPaymentMethodOptions       `json:"bancontact,omitempty"`
	Boleto           *CheckoutBoletoPaymentMethodOptions           `json:"boleto,omitempty"`
	Card             *CheckoutCardPaymentMethodOptions             `json:"card,omitempty"`
	CustomerBalance  *CheckoutCustomerBalancePaymentMethodOptions  `json:"customer_balance,omitempty"`
	Eps              *CheckoutEpsPaymentMethodOptions              `json:"eps,omitempty"`
	Fpx              *CheckoutFpxPaymentMethodOptions              `json:"fpx,omitempty"`
	Giropay          *CheckoutGiropayPaymentMethodOptions          `json:"giropay,omitempty"`
	Grabpay          *CheckoutGrabPayPaymentMethodOptions          `json:"grabpay,omitempty"`
	Ideal            *CheckoutIdealPaymentMethodOptions            `json:"ideal,omitempty"`
	Klarna           *CheckoutKlarnaPaymentMethodOptions           `json:"klarna,omitempty"`
	Konbini          *CheckoutKonbiniPaymentMethodOptions          `json:"konbini,omitempty"`
	Oxxo             *CheckoutOxxoPaymentMethodOptions             `json:"oxxo,omitempty"`
	P24              *CheckoutP24PaymentMethodOptions              `json:"p24,omitempty"`
	Paynow           *CheckoutPaynowPaymentMethodOptions           `json:"paynow,omitempty"`
	Pix              *CheckoutPixPaymentMethodOptions              `json:"pix,omitempty"`
	SepaDebit        *CheckoutSepaDebitPaymentMethodOptions        `json:"sepa_debit,omitempty"`
	Sofort           *CheckoutSofortPaymentMethodOptions           `json:"sofort,omitempty"`
	UsBankAccount    *CheckoutUsBankAccountPaymentMethodOptions    `json:"us_bank_account,omitempty"`
}
