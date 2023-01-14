package attempt

type SetupAttemptPaymentMethodDetails struct {
	AcssDebit     map[string]interface{}                       `json:"acss_debit,omitempty"`
	AuBecsDebit   map[string]interface{}                       `json:"au_becs_debit,omitempty"`
	BacsDebit     map[string]interface{}                       `json:"bacs_debit,omitempty"`
	Bancontact    *SetupAttemptPaymentMethodDetailsBancontact  `json:"bancontact,omitempty"`
	Blik          map[string]interface{}                       `json:"blik,omitempty"`
	Boleto        map[string]interface{}                       `json:"boleto,omitempty"`
	Card          *SetupAttemptPaymentMethodDetailsCard        `json:"card,omitempty"`
	CardPresent   *SetupAttemptPaymentMethodDetailsCardPresent `json:"card_present,omitempty"`
	Ideal         *SetupAttemptPaymentMethodDetailsIdeal       `json:"ideal,omitempty"`
	Klarna        map[string]interface{}                       `json:"klarna,omitempty"`
	Link          map[string]interface{}                       `json:"link,omitempty"`
	SepaDebit     map[string]interface{}                       `json:"sepa_debit,omitempty"`
	Sofort        *SetupAttemptPaymentMethodDetailsSofort      `json:"sofort,omitempty"`
	Type          string                                       `json:"type"`
	UsBankAccount map[string]interface{}                       `json:"us_bank_account,omitempty"`
}
