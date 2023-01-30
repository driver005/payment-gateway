package account

type AccountCapabilities struct {
	AcssDebitPayments          string `json:"acss_debit_payments,omitempty"`
	AffirmPayments             string `json:"affirm_payments,omitempty"`
	AfterpayClearpayPayments   string `json:"afterpay_clearpay_payments,omitempty"`
	AuBecsDebitPayments        string `json:"au_becs_debit_payments,omitempty"`
	BacsDebitPayments          string `json:"bacs_debit_payments,omitempty"`
	BancontactPayments         string `json:"bancontact_payments,omitempty"`
	BankTransferPayments       string `json:"bank_transfer_payments,omitempty"`
	BlikPayments               string `json:"blik_payments,omitempty"`
	BoletoPayments             string `json:"boleto_payments,omitempty"`
	CardIssuing                string `json:"card_issuing,omitempty"`
	CardPayments               string `json:"card_payments,omitempty"`
	CartesBancairesPayments    string `json:"cartes_bancaires_payments,omitempty"`
	EpsPayments                string `json:"eps_payments,omitempty"`
	FpxPayments                string `json:"fpx_payments,omitempty"`
	GiropayPayments            string `json:"giropay_payments,omitempty"`
	GrabpayPayments            string `json:"grabpay_payments,omitempty"`
	IdealPayments              string `json:"ideal_payments,omitempty"`
	IndiaInternationalPayments string `json:"india_international_payments,omitempty"`
	JcbPayments                string `json:"jcb_payments,omitempty"`
	KlarnaPayments             string `json:"klarna_payments,omitempty"`
	KonbiniPayments            string `json:"konbini_payments,omitempty"`
	LegacyPayments             string `json:"legacy_payments,omitempty"`
	LinkPayments               string `json:"link_payments,omitempty"`
	OxxoPayments               string `json:"oxxo_payments,omitempty"`
	P24Payments                string `json:"p24_payments,omitempty"`
	PaynowPayments             string `json:"paynow_payments,omitempty"`
	PromptpayPayments          string `json:"promptpay_payments,omitempty"`
	SepaDebitPayments          string `json:"sepa_debit_payments,omitempty"`
	SofortPayments             string `json:"sofort_payments,omitempty"`
	TaxReportingUs1099K        string `json:"tax_reporting_us_1099_k,omitempty"`
	TaxReportingUs1099Misc     string `json:"tax_reporting_us_1099_misc,omitempty"`
	Transfers                  string `json:"transfers,omitempty"`
	Treasury                   string `json:"treasury,omitempty"`
	UsBankAccountAchPayments   string `json:"us_bank_account_ach_payments,omitempty"`
}
