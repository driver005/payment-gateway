package methods

// PaymentIntentPaymentMethodOptions
type PaymentIntentPaymentMethodOptions struct {
	AcssDebit        *PaymentIntentPaymentMethodOptionsAcssDebit1       `json:"acss_debit,omitempty"`
	Affirm           *PaymentIntentPaymentMethodOptionsAffirm           `json:"affirm,omitempty"`
	AfterpayClearpay *PaymentIntentPaymentMethodOptionsAfterpayClearpay `json:"afterpay_clearpay,omitempty"`
	Alipay           *PaymentIntentPaymentMethodOptionsAlipay           `json:"alipay,omitempty"`
	AuBecsDebit      *PaymentIntentPaymentMethodOptionsAuBecsDebit1     `json:"au_becs_debit,omitempty"`
	BacsDebit        *PaymentIntentPaymentMethodOptionsBacsDebit        `json:"bacs_debit,omitempty"`
	Bancontact       *PaymentIntentPaymentMethodOptionsBancontact       `json:"bancontact,omitempty"`
	Blik             *PaymentIntentPaymentMethodOptionsBlik1            `json:"blik,omitempty"`
	Boleto           *PaymentIntentPaymentMethodOptionsBoleto           `json:"boleto,omitempty"`
	Card             *PaymentIntentPaymentMethodOptionsCard1            `json:"card,omitempty"`
	CardPresent      *PaymentIntentPaymentMethodOptionsCardPresent      `json:"card_present,omitempty"`
	CustomerBalance  *PaymentIntentPaymentMethodOptionsCustomerBalance  `json:"customer_balance,omitempty"`
	Eps              *PaymentIntentPaymentMethodOptionsEps1             `json:"eps,omitempty"`
	Fpx              *PaymentIntentPaymentMethodOptionsFpx              `json:"fpx,omitempty"`
	Giropay          *PaymentIntentPaymentMethodOptionsGiropay          `json:"giropay,omitempty"`
	Grabpay          *PaymentIntentPaymentMethodOptionsGrabpay          `json:"grabpay,omitempty"`
	Ideal            *PaymentIntentPaymentMethodOptionsIdeal            `json:"ideal,omitempty"`
	InteracPresent   *PaymentIntentPaymentMethodOptionsInteracPresent   `json:"interac_present,omitempty"`
	Klarna           *PaymentIntentPaymentMethodOptionsKlarna           `json:"klarna,omitempty"`
	Konbini          *PaymentIntentPaymentMethodOptionsKonbini          `json:"konbini,omitempty"`
	Link             *PaymentIntentPaymentMethodOptionsLink1            `json:"link,omitempty"`
	Oxxo             *PaymentIntentPaymentMethodOptionsOxxo             `json:"oxxo,omitempty"`
	P24              *PaymentIntentPaymentMethodOptionsP24              `json:"p24,omitempty"`
	Paynow           *PaymentIntentPaymentMethodOptionsPaynow           `json:"paynow,omitempty"`
	Pix              *PaymentIntentPaymentMethodOptionsPix              `json:"pix,omitempty"`
	Promptpay        *PaymentIntentPaymentMethodOptionsPromptpay        `json:"promptpay,omitempty"`
	SepaDebit        *PaymentIntentPaymentMethodOptionsSepaDebit1       `json:"sepa_debit,omitempty"`
	Sofort           *PaymentIntentPaymentMethodOptionsSofort           `json:"sofort,omitempty"`
	UsBankAccount    *PaymentIntentPaymentMethodOptionsUsBankAccount1   `json:"us_bank_account,omitempty"`
	WechatPay        *PaymentIntentPaymentMethodOptionsWechatPay        `json:"wechat_pay,omitempty"`
}
