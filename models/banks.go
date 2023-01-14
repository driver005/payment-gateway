package models

import (
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
)

type Account struct{}

type AchCreditTransfer struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	RoutingNumber string `json:"routing_number"`
	SwiftCode     string `json:"brand"`
}

func (a AchCreditTransfer) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AccountNumber),
		validation.Field(&a.BankName),
		validation.Field(&a.RoutingNumber),
		validation.Field(&a.SwiftCode),
	)
}

type AchDebit struct {
	AccountHolderType string `json:"account_holder_type"`
	BankName          string `json:"bank_name"`
	Country           string `json:"country"`
	Fingerprint       string `json:"fingerprint"`
	Last4             string `json:"last4"`
	RoutingNumber     string `json:"routing_number"`
}

func (a AchDebit) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.AccountHolderType, validation.In("individual", "company")),
		validation.Field(&a.BankName),
		validation.Field(&a.Country),
		validation.Field(&a.Fingerprint),
		validation.Field(&a.Last4),
		validation.Field(&a.RoutingNumber),
	)
}

type AcssDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	Mandate           string `json:"mandate"`
	TransitNumber     string `json:"transit_number"`
}

func (a AcssDebit) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.BankName),
		validation.Field(&a.Fingerprint),
		validation.Field(&a.InstitutionNumber),
		validation.Field(&a.Last4),
		validation.Field(&a.Mandate),
		validation.Field(&a.TransitNumber),
	)
}

type AfterpayClearpay struct {
	Reference string `json:"reference"`
}

func (a AfterpayClearpay) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.Reference),
	)
}

type Alipay struct {
	BuyerId       string `json:"buyer_id"`
	Fingerprint   string `json:"fingerprint"`
	TransactionId string `json:"transaction_id"`
}

func (a Alipay) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.BuyerId),
		validation.Field(&a.Fingerprint),
		validation.Field(&a.TransactionId),
	)
}

type AuBecsDebit struct {
	BsbNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

func (a AuBecsDebit) Validate() error {
	return validation.ValidateStruct(&a,
		validation.Field(&a.BsbNumber),
		validation.Field(&a.Fingerprint),
		validation.Field(&a.Last4),
		validation.Field(&a.Mandate),
	)
}

// B
type BacsDebit struct {
	SortCode    string `json:"sort_code"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

func (b BacsDebit) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.Fingerprint),
		validation.Field(&b.Last4),
		validation.Field(&b.Mandate),
		validation.Field(&b.SortCode),
	)
}

type Bankcontact struct {
	BankCode                     string `json:"bank_code"`
	BankName                     string `json:"bank_name"`
	Bic                          string `json:"bic"`
	Generated_sepa_debit         string `json:"generated_sepa_debit"`
	Generated_sepa_debit_mandate string `json:"generated_sepa_debit_mandate"`
	Iban_last4                   string `json:"iban_last4"`
	Preferred_language           string `json:"preferred_language"`
	Verified_name                string `json:"verified_name"`
}

func (b Bankcontact) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.BankCode),
		validation.Field(&b.BankName),
		validation.Field(&b.Bic),
		validation.Field(&b.Generated_sepa_debit),
		validation.Field(&b.Generated_sepa_debit_mandate),
		validation.Field(&b.Iban_last4),
		validation.Field(&b.Preferred_language),
		validation.Field(&b.Verified_name),
	)
}

type Boleto struct {
	TaxId string `json:"tax_id"`
}

func (b Boleto) Validate() error {
	return validation.ValidateStruct(&b,
		validation.Field(&b.TaxId),
	)
}

// C
type CreditCard struct {
	Brand        string       `json:"brand"`
	Checks       Checks       `json:"checks"`
	Country      string       `json:"country"`
	ExpMonth     string       `json:"exp_month"`
	ExpYear      string       `json:"exp_year"`
	Fingerprint  string       `json:"fingerprint"`
	Funding      string       `json:"funding"`
	Installments Installments `json:"installments"`
	Last4        string       `json:"last4"`
	Mandate      string       `json:"mandate"`
	Network      string       `json:"network"`
	ThreeDSecure ThreeDSecure `json:"three_d_secure"`
	Wallet       Wallet       `json:"wallet"`
}

func (c CreditCard) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Brand, validation.In("amex", "diners", "discover", "jcb", "mastercard", "unionpay", "visa", "unknown")),
		validation.Field(&c.Checks.AddressLine1Check, validation.In("pass", "fail", "unavailable", "unchecked")),
		validation.Field(&c.Checks.AddressPostalCodeCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
		validation.Field(&c.Checks.CvcCheck, validation.In("pass", "fail", "unavailable", "unchecked")),
		validation.Field(&c.Country, is.CountryCode2),
		validation.Field(&c.ExpMonth, validation.Min(1), validation.Max(12)),
		validation.Field(&c.ExpYear, validation.Min(2000), validation.Length(4, 4)),
		validation.Field(&c.Fingerprint),
		validation.Field(&c.Funding, validation.In("credit", "debit", "prepaid", "unknown")),
		validation.Field(&c.Installments.Plan.Count),
		validation.Field(&c.Installments.Plan.Interval),
		validation.Field(&c.Installments.Plan.Type),
		validation.Field(&c.Last4),
		validation.Field(&c.Mandate),
		validation.Field(&c.Network, validation.In("amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "visa", "unknown")),
		validation.Field(&c.ThreeDSecure.AuthenticationFlow, validation.In("challenge", "frictionless")),
		validation.Field(&c.ThreeDSecure.Result, validation.In("authenticated", "attempt_acknowledged", "not_supported", "failed", "processing_error")),
		validation.Field(&c.ThreeDSecure.ResultReason, validation.In("card_not_enrolled", "network_not_supported", "abandoned", "canceled", "rejected", "bypassed", "protocol_error")),
		validation.Field(&c.ThreeDSecure.Version, validation.In("challenge", "frictionless")),
		validation.Field(&c.Wallet.AmexExpressCheckout),
		validation.Field(&c.Wallet.ApplePay),
		validation.Field(&c.Wallet.DynamicLast4),
		validation.Field(&c.Wallet.GooglePay),
		validation.Field(&c.Wallet.SamsungPay),
		validation.Field(&c.Wallet.Type, validation.In("amex_express_checkout", "apple_pay", "google_pay", "masterpass", "samsung_pay", "visa_checkout")),
		validation.Field(&c.Wallet.Masterpass.Email, is.Email),
		validation.Field(&c.Wallet.Masterpass.Name),
		validation.Field(&c.Wallet.VisaCheckout.Email, is.Email),
		validation.Field(&c.Wallet.VisaCheckout.Name),
	)
}

type CardPresent struct {
	AmountAuthorized                  int     `json:"amount_authorized"`
	Brand                             string  `json:"brand"`
	CaptureBefore                     string  `json:"capture_before"`
	CardholderName                    string  `json:"cardholder_name"`
	Country                           string  `json:"country"`
	EmvAuthData                       string  `json:"emv_auth_data"`
	ExpMonth                          string  `json:"exp_month"`
	ExpYear                           string  `json:"exp_year"`
	Fingerprint                       string  `json:"fingerprint"`
	Funding                           string  `json:"funding"`
	GeneratedCard                     string  `json:"generated_card"`
	IncrementalAuthorizationSupported string  `json:"incremental_authorization_supported"`
	Last4                             string  `json:"last4"`
	Network                           string  `json:"network"`
	OvercaptureSupported              string  `json:"overcapture_supported"`
	ReadMethod                        string  `json:"read_method"`
	Receipt                           Receipt `json:"receipt"`
}

func (c CardPresent) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.AmountAuthorized),
		validation.Field(&c.Brand, validation.In("amex", "diners", "discover", "jcb", "mastercard", "unionpay", "visa", "unknown")),
		validation.Field(&c.CaptureBefore),
		validation.Field(&c.CardholderName),
		validation.Field(&c.Country, is.CountryCode2),
		validation.Field(&c.ExpMonth, validation.Min(1), validation.Max(12)),
		validation.Field(&c.ExpYear, validation.Min(2000), validation.Length(4, 4)),
		validation.Field(&c.Fingerprint),
		validation.Field(&c.Funding, validation.In("credit", "debit", "prepaid", "unknown")),
		validation.Field(&c.GeneratedCard),
		validation.Field(&c.IncrementalAuthorizationSupported),
		validation.Field(&c.Last4),
		validation.Field(&c.Network, validation.In("amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "visa", "unknown")),
		validation.Field(&c.OvercaptureSupported),
		validation.Field(&c.ReadMethod, validation.In("contact_emv", "contactless_emv", "magnetic_stripe_track2", "magnetic_stripe_fallback", "contactless_magstripe_mode")),
		validation.Field(&c.Receipt.AccountType, validation.In("credit", "checking", "prepaid", "unknown")),
		validation.Field(&c.Receipt.ApplicationCryptogram),
		validation.Field(&c.Receipt.ApplicationPreferredName),
		validation.Field(&c.Receipt.AuthorizationCode),
		validation.Field(&c.Receipt.AuthorizationResponseCode),
		validation.Field(&c.Receipt.CardholderVerificationMethod),
		validation.Field(&c.Receipt.DedicatedFileName),
		validation.Field(&c.Receipt.TerminalVerificationResults),
		validation.Field(&c.Receipt.TransactionStatusInformation),
	)
}

type CustomerBalance struct{}

// E
type Eps struct {
	Bank         string `json:"bank"`
	VerifiedName string `json:"verified_name"`
}

func (e Eps) Validate() error {
	return validation.ValidateStruct(&e,
		validation.Field(&e.Bank, validation.In("arzte_und_apotheker_bank", "austrian_anadi_bank_ag", "bank_austria", "bankhaus_carl_spangler", "bankhaus_schelhammer_und_schattera_ag", "bawag_psk_ag", "bks_bank_ag", "brull_kallmus_bank_ag", "btv_vier_lander_bank", "capital_bank_grawe_gruppe_ag", "dolomitenbank", "easybank_ag", "erste_bank_und_sparkassen", "hypo_alpeadriabank_international_ag", "hypo_noe_lb_fur_niederosterreich_u_wien", "hypo_oberosterreich_salzburg_steiermark", "hypo_tirol_bank_ag", "hypo_vorarlberg_bank_ag", "hypo_bank_burgenland_aktiengesellschaft", "marchfelder_bank", "oberbank_ag", "raiffeisen_bankengruppe_osterreich", "schoellerbank_ag", "sparda_bank_wien", "volksbank_gruppe", "volkskreditbank_ag", "vr_bank_braunau")),
		validation.Field(&e.VerifiedName),
	)
}

// F
type Fpx struct {
	Bank          string `json:"bank"`
	TransactionId string `json:"transaction_id"`
}

func (f Fpx) Validate() error {
	return validation.ValidateStruct(&f,
		validation.Field(&f.Bank, validation.In("affin_bank", "agrobank", "alliance_bank", "ambank", "bank_islam", "bank_muamalat", "bank_rakyat", "bsn", "cimb", "hong_leong_bank", "hsbc", "kfh", "maybank2u", "ocbc", "public_bank", "rhb", "standard_chartered", "uob", "deutsche_bank", "maybank2e", "pb_enterprise")),
		validation.Field(&f.TransactionId),
	)
}

// G
type Giropay struct {
	BankCode     string `json:"bank_code"`
	BankName     string `json:"bank_name"`
	Bic          string `json:"bic"`
	VerifiedName string `json:"verified_name"`
}

func (g Giropay) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.BankCode),
		validation.Field(&g.BankName),
		validation.Field(&g.Bic),
		validation.Field(&g.VerifiedName),
	)
}

type Grabpay struct {
	TransactionId string `json:"transaction_id"`
}

func (g Grabpay) Validate() error {
	return validation.ValidateStruct(&g,
		validation.Field(&g.TransactionId),
	)
}

// I
type Ideal struct {
	Bank                      string `json:"bank"`
	Bic                       string `json:"bic"`
	GeneratedSepaDebit        string `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate string `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string `json:"iban_last4"`
	VerifiedName              string `json:"verified_name"`
}

func (i Ideal) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Bank, validation.In("abn_amro", "asn_bank", "bunq", "handelsbanken", "ing", "knab", "moneyou", "rabobank", "regiobank", "revolut", "sns_bank", "triodos_bank", "van_lanschot")),
		validation.Field(&i.Bic),
		validation.Field(&i.GeneratedSepaDebit),
		validation.Field(&i.GeneratedSepaDebitMandate),
		validation.Field(&i.IbanLast4),
		validation.Field(&i.VerifiedName),
	)
}

type InteracPresent struct {
	Brand            string  `json:"brand"`
	CardholderName   string  `json:"cardholder_name"`
	Country          string  `json:"country"`
	EmvAuthData      string  `json:"emv_auth_data"`
	ExpMonth         string  `json:"exp_month"`
	ExpYear          string  `json:"exp_year"`
	Fingerprint      string  `json:"fingerprint"`
	Funding          string  `json:"funding"`
	GeneratedCard    string  `json:"generated_card"`
	Last4            string  `json:"last4"`
	Network          string  `json:"network"`
	PreferredLocales string  `json:"preferred_locales"`
	ReadMethod       string  `json:"read_method"`
	Receipt          Receipt `json:"receipt"`
}

func (i InteracPresent) Validate() error {
	return validation.ValidateStruct(&i,
		validation.Field(&i.Brand, validation.In("interac", "mastercard", "visa")),
		validation.Field(&i.EmvAuthData),
		validation.Field(&i.CardholderName),
		validation.Field(&i.Country, is.CountryCode2),
		validation.Field(&i.ExpMonth, validation.Min(1), validation.Max(12)),
		validation.Field(&i.ExpYear, validation.Min(2000), validation.Length(4, 4)),
		validation.Field(&i.Fingerprint),
		validation.Field(&i.Funding, validation.In("credit", "debit", "prepaid", "unknown")),
		validation.Field(&i.GeneratedCard),
		validation.Field(&i.Last4),
		validation.Field(&i.Network, validation.In("amex", "cartes_bancaires", "diners", "discover", "interac", "jcb", "mastercard", "unionpay", "visa", "unknown")),
		validation.Field(&i.PreferredLocales),
		validation.Field(&i.ReadMethod, validation.In("contact_emv", "contactless_emv", "magnetic_stripe_track2", "magnetic_stripe_fallback", "contactless_magstripe_mode")),
		validation.Field(&i.Receipt.AccountType, validation.In("checking", "savings", "unknown")),
		validation.Field(&i.Receipt.ApplicationCryptogram),
		validation.Field(&i.Receipt.ApplicationPreferredName),
		validation.Field(&i.Receipt.AuthorizationCode),
		validation.Field(&i.Receipt.AuthorizationResponseCode),
		validation.Field(&i.Receipt.CardholderVerificationMethod),
		validation.Field(&i.Receipt.DedicatedFileName),
		validation.Field(&i.Receipt.TerminalVerificationResults),
		validation.Field(&i.Receipt.TransactionStatusInformation),
	)
}

// K
type Klarna struct {
	PaymentMethodCategory string `json:"payment_method_category"`
	PreferredLocale       string `json:"preferred_locale"`
}

func (k Klarna) Validate() error {
	return validation.ValidateStruct(&k,
		validation.Field(&k.PaymentMethodCategory, validation.In("pay_later", "pay_now", "pay_with_financing", "pay_in_installments")),
		validation.Field(&k.PreferredLocale, validation.In("de-AT", "en-AT", "nl-BE", "fr-BE", "en-BE", "de-DE", "en-DE", "da-DK", "en-DK", "es-ES", "en-ES", "fi-FI", "sv-FI", "en-FI", "en-GB", "en-IE", "it-IT", "en-IT", "nl-NL", "en-NL", "nb-NO", "en-NO", "sv-SE", "en-SE", "en-US", "es-US", "fr-FR", "en-FR")),
	)
}

type Konbini struct {
	Store Store `json:"store"`
}

func (k Konbini) Validate() error {
	return validation.ValidateStruct(&k,
		validation.Field(&k.Store.Chain),
	)
}

// M
type Multibanco struct {
	Entity    string `json:"entity"`
	Reference string `json:"reference"`
}

func (m Multibanco) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.Entity),
		validation.Field(&m.Reference),
	)
}

// O
type Oxxo struct {
	Number string `json:"number"`
}

func (o Oxxo) Validate() error {
	return validation.ValidateStruct(&o,
		validation.Field(&o.Number),
	)
}

// P
type P24 struct {
	Bank         string `json:"bank"`
	Reference    string `json:"reference"`
	VerifiedName string `json:"verified_name"`
}

func (p P24) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Bank, validation.In("ing", "citi_handlowy", "tmobile_usbugi_bankowe", "plus_bank", "etransfer_pocztowy24", "banki_spbdzielcze", "bank_nowy_bfg_sa", "getin_bank", "blik", "noble_pay", "ideabank", "envelobank", "santander_przelew24", "nest_przelew", "mbank_mtransfer", "inteligo", "pbac_z_ipko", "bnp_paribas", "credit_agricole", "toyota_bank", "bank_pekao_sa", "volkswagen_bank", "bank_millennium", "alior_bank", "boz")),
		validation.Field(&p.Reference),
		validation.Field(&p.VerifiedName),
	)
}

type Paynow struct {
	Reference string `json:"reference"`
}

func (p Paynow) Validate() error {
	return validation.ValidateStruct(&p,
		validation.Field(&p.Reference),
	)
}

// S
type SepaDebit struct {
	BankCode    string `json:"bank_code"`
	BranchCode  string `json:"branch_code"`
	Country     string `json:"country"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}

func (s SepaDebit) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.BankCode),
		validation.Field(&s.BranchCode),
		validation.Field(&s.Country, is.CountryCode2),
		validation.Field(&s.Fingerprint),
		validation.Field(&s.Last4),
		validation.Field(&s.Mandate),
	)
}

type Sofort struct {
	BankCode                  string `json:"bank_code"`
	BankName                  string `json:"bank_name"`
	Bic                       string `json:"bic"`
	Country                   string `json:"country"`
	GeneratedsepaDebit        string `json:"generated_sepa_debit"`
	GeneratedsepaDebitMandate string `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string `json:"iban_last4"`
	PreferredLanguage         string `json:"preferred_language"`
	VerifiedName              string `json:"verified_name"`
}

func (s Sofort) Validate() error {
	return validation.ValidateStruct(&s,
		validation.Field(&s.BankCode),
		validation.Field(&s.BankName),
		validation.Field(&s.Bic),
		validation.Field(&s.Country, is.CountryCode2),
		validation.Field(&s.GeneratedsepaDebit),
		validation.Field(&s.GeneratedsepaDebitMandate),
		validation.Field(&s.IbanLast4),
		validation.Field(&s.PreferredLanguage),
		validation.Field(&s.VerifiedName),
	)
}

// U
type UsBankAccount struct {
	AccountHolderType string `json:"account_holder_type"`
	AccountType       string `json:"account_type"`
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	Last4             string `json:"last4"`
	RoutingNumber     string `json:"routing_number"`
}

func (u UsBankAccount) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.AccountHolderType, validation.In("individual", "company")),
		validation.Field(&u.AccountType, validation.In("checking", "savings")),
		validation.Field(&u.BankName),
		validation.Field(&u.Fingerprint),
		validation.Field(&u.Last4),
		validation.Field(&u.RoutingNumber),
	)
}

// W
type Wechat struct{}

type WechatPay struct {
	Fingerprint   string `json:"fingerprint"`
	TransactionId string `json:"transaction_id"`
}

func (w WechatPay) Validate() error {
	return validation.ValidateStruct(&w,
		validation.Field(&w.Fingerprint),
		validation.Field(&w.TransactionId),
	)
}
