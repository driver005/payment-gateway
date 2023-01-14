package methods

type MandatePaymentMethodDetails struct {
	AcssDebit   *MandateAcssDebit   `json:"acss_debit,omitempty"`
	AuBecsDebit *MandateAuBecsDebit `json:"au_becs_debit,omitempty"`
	BacsDebit   *MandateBacsDebit   `json:"bacs_debit,omitempty"`
	Blik        *MandateBlik        `json:"blik,omitempty"`
	Card *MandateCard `json:"card,omitempty"`
	Link      *MandateLink `json:"link,omitempty"`
	SepaDebit *MandateSepaDebit      `json:"sepa_debit,omitempty"`
	UsBankAccount *MandateUsBankAccount `json:"us_bank_account,omitempty"`
	Type string `json:"type"`
}
