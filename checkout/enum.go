package checkout

import "database/sql/driver"

type Mode string

const (
	ModePayment      Mode = "payment"
	ModeSetup        Mode = "setup"
	ModeSubscription Mode = "subscription"
)

func (ct *Mode) Scan(value interface{}) error {
	*ct = Mode(value.(string))
	return nil
}

func (ct Mode) Value() (driver.Value, error) {
	return string(ct), nil
}

type PaymentStatus string

const (
	PaymentStatusPaid              PaymentStatus = "paid"
	PaymentStatusUnpaid            PaymentStatus = "unpaid"
	PaymentStatusNoPaymentRequired PaymentStatus = "no_payment_required"
)

func (ct *PaymentStatus) Scan(value interface{}) error {
	*ct = PaymentStatus(value.(string))
	return nil
}

func (ct PaymentStatus) Value() (driver.Value, error) {
	return string(ct), nil
}

type Status string

const (
	StatusOpen     Status = "open"
	StatusComplete Status = "complete"
	StatusExpired  Status = "expired"
)

func (ct *Status) Scan(value interface{}) error {
	*ct = Status(value.(string))
	return nil
}

func (ct Status) Value() (driver.Value, error) {
	return string(ct), nil
}

type BillingAddressCollection string

const (
	BillingAddressCollectionAuto     BillingAddressCollection = "auto"
	BillingAddressCollectionRequired BillingAddressCollection = "required"
)

func (ct *BillingAddressCollection) Scan(value interface{}) error {
	*ct = BillingAddressCollection(value.(string))
	return nil
}

func (ct BillingAddressCollection) Value() (driver.Value, error) {
	return string(ct), nil
}

type CustomerCreation string

const (
	CustomerCreationIfRequired CustomerCreation = "if_required"
	CustomerCreationAlways     CustomerCreation = "always"
)

func (ct *CustomerCreation) Scan(value interface{}) error {
	*ct = CustomerCreation(value.(string))
	return nil
}

func (ct CustomerCreation) Value() (driver.Value, error) {
	return string(ct), nil
}

type Locale string

const (
	Localeauto   Locale = "auto"
	Localebg     Locale = "bg"
	Localecs     Locale = "cs"
	Localeda     Locale = "da"
	Localede     Locale = "de"
	Localeel     Locale = "el"
	Localeen     Locale = "en"
	Localeen_GB  Locale = "en-GB"
	Localees     Locale = "es"
	Localees_419 Locale = "es-419"
	Localeet     Locale = "et"
	Localefi     Locale = "fi"
	Localefil    Locale = "fil"
	Localefr     Locale = "fr"
	Localefr_CA  Locale = "fr-CA"
	Localehr     Locale = "hr"
	Localehu     Locale = "hu"
	Localeid     Locale = "id"
	Localeit     Locale = "it"
	Localeja     Locale = "ja"
	Localeko     Locale = "ko"
	Localelt     Locale = "lt"
	Localelv     Locale = "lv"
	Localems     Locale = "ms"
	Localemt     Locale = "mt"
	Localenb     Locale = "nb"
	Localenl     Locale = "nl"
	Localepl     Locale = "pl"
	Localept     Locale = "pt"
	Localept_BR  Locale = "pt-BR"
	Localero     Locale = "ro"
	Localeru     Locale = "ru"
	Localesk     Locale = "sk"
	Localesl     Locale = "sl"
	Localesv     Locale = "sv"
	Localeth     Locale = "th"
	Localetr     Locale = "tr"
	Localevi     Locale = "vi"
	Localezh     Locale = "zh"
	Localezh_HK  Locale = "zh-HK"
	Localezh_TW  Locale = "zh-TW"
)

func (ct *Locale) Scan(value interface{}) error {
	*ct = Locale(value.(string))
	return nil
}

func (ct Locale) Value() (driver.Value, error) {
	return string(ct), nil
}

type PaymentMethodCollection string

const (
	PaymentMethodCollectionAlways     PaymentMethodCollection = "always"
	PaymentMethodCollectionIfRequired PaymentMethodCollection = "if_required"
)

func (ct *PaymentMethodCollection) Scan(value interface{}) error {
	*ct = PaymentMethodCollection(value.(string))
	return nil
}

func (ct PaymentMethodCollection) Value() (driver.Value, error) {
	return string(ct), nil
}

type SubmitType string

const (
	SubmitTypeAuto   SubmitType = "auto"
	SubmitTypePay    SubmitType = "pay"
	SubmitTypeBook   SubmitType = "book"
	SubmitTypeDonate SubmitType = "donate"
)

func (ct *SubmitType) Scan(value interface{}) error {
	*ct = SubmitType(value.(string))
	return nil
}

func (ct SubmitType) Value() (driver.Value, error) {
	return string(ct), nil
}
