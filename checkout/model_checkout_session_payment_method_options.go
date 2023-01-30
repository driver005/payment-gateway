package checkout

import (
	"encoding/json"
)

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

// NewCheckoutSessionPaymentMethodOptions instantiates a new CheckoutSessionPaymentMethodOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutSessionPaymentMethodOptions() *CheckoutSessionPaymentMethodOptions {
	this := CheckoutSessionPaymentMethodOptions{}
	return &this
}

// NewCheckoutSessionPaymentMethodOptionsWithDefaults instantiates a new CheckoutSessionPaymentMethodOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutSessionPaymentMethodOptionsWithDefaults() *CheckoutSessionPaymentMethodOptions {
	this := CheckoutSessionPaymentMethodOptions{}
	return &this
}

// GetAcssDebit returns the AcssDebit field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetAcssDebit() CheckoutAcssDebitPaymentMethodOptions {
	if o == nil || isNil(o.AcssDebit) {
		var ret CheckoutAcssDebitPaymentMethodOptions
		return ret
	}
	return *o.AcssDebit
}

// GetAcssDebitOk returns a tuple with the AcssDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetAcssDebitOk() (*CheckoutAcssDebitPaymentMethodOptions, bool) {
	if o == nil || isNil(o.AcssDebit) {
		return nil, false
	}
	return o.AcssDebit, true
}

// HasAcssDebit returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasAcssDebit() bool {
	if o != nil && !isNil(o.AcssDebit) {
		return true
	}

	return false
}

// SetAcssDebit gets a reference to the given CheckoutAcssDebitPaymentMethodOptions and assigns it to the AcssDebit field.
func (o *CheckoutSessionPaymentMethodOptions) SetAcssDebit(v CheckoutAcssDebitPaymentMethodOptions) {
	o.AcssDebit = &v
}

// GetAffirm returns the Affirm field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetAffirm() CheckoutAffirmPaymentMethodOptions {
	if o == nil || isNil(o.Affirm) {
		var ret CheckoutAffirmPaymentMethodOptions
		return ret
	}
	return *o.Affirm
}

// GetAffirmOk returns a tuple with the Affirm field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetAffirmOk() (*CheckoutAffirmPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Affirm) {
		return nil, false
	}
	return o.Affirm, true
}

// HasAffirm returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasAffirm() bool {
	if o != nil && !isNil(o.Affirm) {
		return true
	}

	return false
}

// SetAffirm gets a reference to the given CheckoutAffirmPaymentMethodOptions and assigns it to the Affirm field.
func (o *CheckoutSessionPaymentMethodOptions) SetAffirm(v CheckoutAffirmPaymentMethodOptions) {
	o.Affirm = &v
}

// GetAfterpayClearpay returns the AfterpayClearpay field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetAfterpayClearpay() CheckoutAfterpayClearpayPaymentMethodOptions {
	if o == nil || isNil(o.AfterpayClearpay) {
		var ret CheckoutAfterpayClearpayPaymentMethodOptions
		return ret
	}
	return *o.AfterpayClearpay
}

// GetAfterpayClearpayOk returns a tuple with the AfterpayClearpay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetAfterpayClearpayOk() (*CheckoutAfterpayClearpayPaymentMethodOptions, bool) {
	if o == nil || isNil(o.AfterpayClearpay) {
		return nil, false
	}
	return o.AfterpayClearpay, true
}

// HasAfterpayClearpay returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasAfterpayClearpay() bool {
	if o != nil && !isNil(o.AfterpayClearpay) {
		return true
	}

	return false
}

// SetAfterpayClearpay gets a reference to the given CheckoutAfterpayClearpayPaymentMethodOptions and assigns it to the AfterpayClearpay field.
func (o *CheckoutSessionPaymentMethodOptions) SetAfterpayClearpay(v CheckoutAfterpayClearpayPaymentMethodOptions) {
	o.AfterpayClearpay = &v
}

// GetAlipay returns the Alipay field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetAlipay() CheckoutAlipayPaymentMethodOptions {
	if o == nil || isNil(o.Alipay) {
		var ret CheckoutAlipayPaymentMethodOptions
		return ret
	}
	return *o.Alipay
}

// GetAlipayOk returns a tuple with the Alipay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetAlipayOk() (*CheckoutAlipayPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Alipay) {
		return nil, false
	}
	return o.Alipay, true
}

// HasAlipay returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasAlipay() bool {
	if o != nil && !isNil(o.Alipay) {
		return true
	}

	return false
}

// SetAlipay gets a reference to the given CheckoutAlipayPaymentMethodOptions and assigns it to the Alipay field.
func (o *CheckoutSessionPaymentMethodOptions) SetAlipay(v CheckoutAlipayPaymentMethodOptions) {
	o.Alipay = &v
}

// GetAuBecsDebit returns the AuBecsDebit field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetAuBecsDebit() CheckoutAuBecsDebitPaymentMethodOptions {
	if o == nil || isNil(o.AuBecsDebit) {
		var ret CheckoutAuBecsDebitPaymentMethodOptions
		return ret
	}
	return *o.AuBecsDebit
}

// GetAuBecsDebitOk returns a tuple with the AuBecsDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetAuBecsDebitOk() (*CheckoutAuBecsDebitPaymentMethodOptions, bool) {
	if o == nil || isNil(o.AuBecsDebit) {
		return nil, false
	}
	return o.AuBecsDebit, true
}

// HasAuBecsDebit returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasAuBecsDebit() bool {
	if o != nil && !isNil(o.AuBecsDebit) {
		return true
	}

	return false
}

// SetAuBecsDebit gets a reference to the given CheckoutAuBecsDebitPaymentMethodOptions and assigns it to the AuBecsDebit field.
func (o *CheckoutSessionPaymentMethodOptions) SetAuBecsDebit(v CheckoutAuBecsDebitPaymentMethodOptions) {
	o.AuBecsDebit = &v
}

// GetBacsDebit returns the BacsDebit field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetBacsDebit() CheckoutBacsDebitPaymentMethodOptions {
	if o == nil || isNil(o.BacsDebit) {
		var ret CheckoutBacsDebitPaymentMethodOptions
		return ret
	}
	return *o.BacsDebit
}

// GetBacsDebitOk returns a tuple with the BacsDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetBacsDebitOk() (*CheckoutBacsDebitPaymentMethodOptions, bool) {
	if o == nil || isNil(o.BacsDebit) {
		return nil, false
	}
	return o.BacsDebit, true
}

// HasBacsDebit returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasBacsDebit() bool {
	if o != nil && !isNil(o.BacsDebit) {
		return true
	}

	return false
}

// SetBacsDebit gets a reference to the given CheckoutBacsDebitPaymentMethodOptions and assigns it to the BacsDebit field.
func (o *CheckoutSessionPaymentMethodOptions) SetBacsDebit(v CheckoutBacsDebitPaymentMethodOptions) {
	o.BacsDebit = &v
}

// GetBancontact returns the Bancontact field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetBancontact() CheckoutBancontactPaymentMethodOptions {
	if o == nil || isNil(o.Bancontact) {
		var ret CheckoutBancontactPaymentMethodOptions
		return ret
	}
	return *o.Bancontact
}

// GetBancontactOk returns a tuple with the Bancontact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetBancontactOk() (*CheckoutBancontactPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Bancontact) {
		return nil, false
	}
	return o.Bancontact, true
}

// HasBancontact returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasBancontact() bool {
	if o != nil && !isNil(o.Bancontact) {
		return true
	}

	return false
}

// SetBancontact gets a reference to the given CheckoutBancontactPaymentMethodOptions and assigns it to the Bancontact field.
func (o *CheckoutSessionPaymentMethodOptions) SetBancontact(v CheckoutBancontactPaymentMethodOptions) {
	o.Bancontact = &v
}

// GetBoleto returns the Boleto field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetBoleto() CheckoutBoletoPaymentMethodOptions {
	if o == nil || isNil(o.Boleto) {
		var ret CheckoutBoletoPaymentMethodOptions
		return ret
	}
	return *o.Boleto
}

// GetBoletoOk returns a tuple with the Boleto field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetBoletoOk() (*CheckoutBoletoPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Boleto) {
		return nil, false
	}
	return o.Boleto, true
}

// HasBoleto returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasBoleto() bool {
	if o != nil && !isNil(o.Boleto) {
		return true
	}

	return false
}

// SetBoleto gets a reference to the given CheckoutBoletoPaymentMethodOptions and assigns it to the Boleto field.
func (o *CheckoutSessionPaymentMethodOptions) SetBoleto(v CheckoutBoletoPaymentMethodOptions) {
	o.Boleto = &v
}

// GetCard returns the Card field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetCard() CheckoutCardPaymentMethodOptions {
	if o == nil || isNil(o.Card) {
		var ret CheckoutCardPaymentMethodOptions
		return ret
	}
	return *o.Card
}

// GetCardOk returns a tuple with the Card field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetCardOk() (*CheckoutCardPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Card) {
		return nil, false
	}
	return o.Card, true
}

// HasCard returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasCard() bool {
	if o != nil && !isNil(o.Card) {
		return true
	}

	return false
}

// SetCard gets a reference to the given CheckoutCardPaymentMethodOptions and assigns it to the Card field.
func (o *CheckoutSessionPaymentMethodOptions) SetCard(v CheckoutCardPaymentMethodOptions) {
	o.Card = &v
}

// GetCustomerBalance returns the CustomerBalance field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetCustomerBalance() CheckoutCustomerBalancePaymentMethodOptions {
	if o == nil || isNil(o.CustomerBalance) {
		var ret CheckoutCustomerBalancePaymentMethodOptions
		return ret
	}
	return *o.CustomerBalance
}

// GetCustomerBalanceOk returns a tuple with the CustomerBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetCustomerBalanceOk() (*CheckoutCustomerBalancePaymentMethodOptions, bool) {
	if o == nil || isNil(o.CustomerBalance) {
		return nil, false
	}
	return o.CustomerBalance, true
}

// HasCustomerBalance returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasCustomerBalance() bool {
	if o != nil && !isNil(o.CustomerBalance) {
		return true
	}

	return false
}

// SetCustomerBalance gets a reference to the given CheckoutCustomerBalancePaymentMethodOptions and assigns it to the CustomerBalance field.
func (o *CheckoutSessionPaymentMethodOptions) SetCustomerBalance(v CheckoutCustomerBalancePaymentMethodOptions) {
	o.CustomerBalance = &v
}

// GetEps returns the Eps field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetEps() CheckoutEpsPaymentMethodOptions {
	if o == nil || isNil(o.Eps) {
		var ret CheckoutEpsPaymentMethodOptions
		return ret
	}
	return *o.Eps
}

// GetEpsOk returns a tuple with the Eps field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetEpsOk() (*CheckoutEpsPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Eps) {
		return nil, false
	}
	return o.Eps, true
}

// HasEps returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasEps() bool {
	if o != nil && !isNil(o.Eps) {
		return true
	}

	return false
}

// SetEps gets a reference to the given CheckoutEpsPaymentMethodOptions and assigns it to the Eps field.
func (o *CheckoutSessionPaymentMethodOptions) SetEps(v CheckoutEpsPaymentMethodOptions) {
	o.Eps = &v
}

// GetFpx returns the Fpx field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetFpx() CheckoutFpxPaymentMethodOptions {
	if o == nil || isNil(o.Fpx) {
		var ret CheckoutFpxPaymentMethodOptions
		return ret
	}
	return *o.Fpx
}

// GetFpxOk returns a tuple with the Fpx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetFpxOk() (*CheckoutFpxPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Fpx) {
		return nil, false
	}
	return o.Fpx, true
}

// HasFpx returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasFpx() bool {
	if o != nil && !isNil(o.Fpx) {
		return true
	}

	return false
}

// SetFpx gets a reference to the given CheckoutFpxPaymentMethodOptions and assigns it to the Fpx field.
func (o *CheckoutSessionPaymentMethodOptions) SetFpx(v CheckoutFpxPaymentMethodOptions) {
	o.Fpx = &v
}

// GetGiropay returns the Giropay field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetGiropay() CheckoutGiropayPaymentMethodOptions {
	if o == nil || isNil(o.Giropay) {
		var ret CheckoutGiropayPaymentMethodOptions
		return ret
	}
	return *o.Giropay
}

// GetGiropayOk returns a tuple with the Giropay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetGiropayOk() (*CheckoutGiropayPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Giropay) {
		return nil, false
	}
	return o.Giropay, true
}

// HasGiropay returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasGiropay() bool {
	if o != nil && !isNil(o.Giropay) {
		return true
	}

	return false
}

// SetGiropay gets a reference to the given CheckoutGiropayPaymentMethodOptions and assigns it to the Giropay field.
func (o *CheckoutSessionPaymentMethodOptions) SetGiropay(v CheckoutGiropayPaymentMethodOptions) {
	o.Giropay = &v
}

// GetGrabpay returns the Grabpay field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetGrabpay() CheckoutGrabPayPaymentMethodOptions {
	if o == nil || isNil(o.Grabpay) {
		var ret CheckoutGrabPayPaymentMethodOptions
		return ret
	}
	return *o.Grabpay
}

// GetGrabpayOk returns a tuple with the Grabpay field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetGrabpayOk() (*CheckoutGrabPayPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Grabpay) {
		return nil, false
	}
	return o.Grabpay, true
}

// HasGrabpay returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasGrabpay() bool {
	if o != nil && !isNil(o.Grabpay) {
		return true
	}

	return false
}

// SetGrabpay gets a reference to the given CheckoutGrabPayPaymentMethodOptions and assigns it to the Grabpay field.
func (o *CheckoutSessionPaymentMethodOptions) SetGrabpay(v CheckoutGrabPayPaymentMethodOptions) {
	o.Grabpay = &v
}

// GetIdeal returns the Ideal field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetIdeal() CheckoutIdealPaymentMethodOptions {
	if o == nil || isNil(o.Ideal) {
		var ret CheckoutIdealPaymentMethodOptions
		return ret
	}
	return *o.Ideal
}

// GetIdealOk returns a tuple with the Ideal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetIdealOk() (*CheckoutIdealPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Ideal) {
		return nil, false
	}
	return o.Ideal, true
}

// HasIdeal returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasIdeal() bool {
	if o != nil && !isNil(o.Ideal) {
		return true
	}

	return false
}

// SetIdeal gets a reference to the given CheckoutIdealPaymentMethodOptions and assigns it to the Ideal field.
func (o *CheckoutSessionPaymentMethodOptions) SetIdeal(v CheckoutIdealPaymentMethodOptions) {
	o.Ideal = &v
}

// GetKlarna returns the Klarna field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetKlarna() CheckoutKlarnaPaymentMethodOptions {
	if o == nil || isNil(o.Klarna) {
		var ret CheckoutKlarnaPaymentMethodOptions
		return ret
	}
	return *o.Klarna
}

// GetKlarnaOk returns a tuple with the Klarna field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetKlarnaOk() (*CheckoutKlarnaPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Klarna) {
		return nil, false
	}
	return o.Klarna, true
}

// HasKlarna returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasKlarna() bool {
	if o != nil && !isNil(o.Klarna) {
		return true
	}

	return false
}

// SetKlarna gets a reference to the given CheckoutKlarnaPaymentMethodOptions and assigns it to the Klarna field.
func (o *CheckoutSessionPaymentMethodOptions) SetKlarna(v CheckoutKlarnaPaymentMethodOptions) {
	o.Klarna = &v
}

// GetKonbini returns the Konbini field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetKonbini() CheckoutKonbiniPaymentMethodOptions {
	if o == nil || isNil(o.Konbini) {
		var ret CheckoutKonbiniPaymentMethodOptions
		return ret
	}
	return *o.Konbini
}

// GetKonbiniOk returns a tuple with the Konbini field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetKonbiniOk() (*CheckoutKonbiniPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Konbini) {
		return nil, false
	}
	return o.Konbini, true
}

// HasKonbini returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasKonbini() bool {
	if o != nil && !isNil(o.Konbini) {
		return true
	}

	return false
}

// SetKonbini gets a reference to the given CheckoutKonbiniPaymentMethodOptions and assigns it to the Konbini field.
func (o *CheckoutSessionPaymentMethodOptions) SetKonbini(v CheckoutKonbiniPaymentMethodOptions) {
	o.Konbini = &v
}

// GetOxxo returns the Oxxo field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetOxxo() CheckoutOxxoPaymentMethodOptions {
	if o == nil || isNil(o.Oxxo) {
		var ret CheckoutOxxoPaymentMethodOptions
		return ret
	}
	return *o.Oxxo
}

// GetOxxoOk returns a tuple with the Oxxo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetOxxoOk() (*CheckoutOxxoPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Oxxo) {
		return nil, false
	}
	return o.Oxxo, true
}

// HasOxxo returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasOxxo() bool {
	if o != nil && !isNil(o.Oxxo) {
		return true
	}

	return false
}

// SetOxxo gets a reference to the given CheckoutOxxoPaymentMethodOptions and assigns it to the Oxxo field.
func (o *CheckoutSessionPaymentMethodOptions) SetOxxo(v CheckoutOxxoPaymentMethodOptions) {
	o.Oxxo = &v
}

// GetP24 returns the P24 field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetP24() CheckoutP24PaymentMethodOptions {
	if o == nil || isNil(o.P24) {
		var ret CheckoutP24PaymentMethodOptions
		return ret
	}
	return *o.P24
}

// GetP24Ok returns a tuple with the P24 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetP24Ok() (*CheckoutP24PaymentMethodOptions, bool) {
	if o == nil || isNil(o.P24) {
		return nil, false
	}
	return o.P24, true
}

// HasP24 returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasP24() bool {
	if o != nil && !isNil(o.P24) {
		return true
	}

	return false
}

// SetP24 gets a reference to the given CheckoutP24PaymentMethodOptions and assigns it to the P24 field.
func (o *CheckoutSessionPaymentMethodOptions) SetP24(v CheckoutP24PaymentMethodOptions) {
	o.P24 = &v
}

// GetPaynow returns the Paynow field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetPaynow() CheckoutPaynowPaymentMethodOptions {
	if o == nil || isNil(o.Paynow) {
		var ret CheckoutPaynowPaymentMethodOptions
		return ret
	}
	return *o.Paynow
}

// GetPaynowOk returns a tuple with the Paynow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetPaynowOk() (*CheckoutPaynowPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Paynow) {
		return nil, false
	}
	return o.Paynow, true
}

// HasPaynow returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasPaynow() bool {
	if o != nil && !isNil(o.Paynow) {
		return true
	}

	return false
}

// SetPaynow gets a reference to the given CheckoutPaynowPaymentMethodOptions and assigns it to the Paynow field.
func (o *CheckoutSessionPaymentMethodOptions) SetPaynow(v CheckoutPaynowPaymentMethodOptions) {
	o.Paynow = &v
}

// GetPix returns the Pix field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetPix() CheckoutPixPaymentMethodOptions {
	if o == nil || isNil(o.Pix) {
		var ret CheckoutPixPaymentMethodOptions
		return ret
	}
	return *o.Pix
}

// GetPixOk returns a tuple with the Pix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetPixOk() (*CheckoutPixPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Pix) {
		return nil, false
	}
	return o.Pix, true
}

// HasPix returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasPix() bool {
	if o != nil && !isNil(o.Pix) {
		return true
	}

	return false
}

// SetPix gets a reference to the given CheckoutPixPaymentMethodOptions and assigns it to the Pix field.
func (o *CheckoutSessionPaymentMethodOptions) SetPix(v CheckoutPixPaymentMethodOptions) {
	o.Pix = &v
}

// GetSepaDebit returns the SepaDebit field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetSepaDebit() CheckoutSepaDebitPaymentMethodOptions {
	if o == nil || isNil(o.SepaDebit) {
		var ret CheckoutSepaDebitPaymentMethodOptions
		return ret
	}
	return *o.SepaDebit
}

// GetSepaDebitOk returns a tuple with the SepaDebit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetSepaDebitOk() (*CheckoutSepaDebitPaymentMethodOptions, bool) {
	if o == nil || isNil(o.SepaDebit) {
		return nil, false
	}
	return o.SepaDebit, true
}

// HasSepaDebit returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasSepaDebit() bool {
	if o != nil && !isNil(o.SepaDebit) {
		return true
	}

	return false
}

// SetSepaDebit gets a reference to the given CheckoutSepaDebitPaymentMethodOptions and assigns it to the SepaDebit field.
func (o *CheckoutSessionPaymentMethodOptions) SetSepaDebit(v CheckoutSepaDebitPaymentMethodOptions) {
	o.SepaDebit = &v
}

// GetSofort returns the Sofort field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetSofort() CheckoutSofortPaymentMethodOptions {
	if o == nil || isNil(o.Sofort) {
		var ret CheckoutSofortPaymentMethodOptions
		return ret
	}
	return *o.Sofort
}

// GetSofortOk returns a tuple with the Sofort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetSofortOk() (*CheckoutSofortPaymentMethodOptions, bool) {
	if o == nil || isNil(o.Sofort) {
		return nil, false
	}
	return o.Sofort, true
}

// HasSofort returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasSofort() bool {
	if o != nil && !isNil(o.Sofort) {
		return true
	}

	return false
}

// SetSofort gets a reference to the given CheckoutSofortPaymentMethodOptions and assigns it to the Sofort field.
func (o *CheckoutSessionPaymentMethodOptions) SetSofort(v CheckoutSofortPaymentMethodOptions) {
	o.Sofort = &v
}

// GetUsBankAccount returns the UsBankAccount field value if set, zero value otherwise.
func (o *CheckoutSessionPaymentMethodOptions) GetUsBankAccount() CheckoutUsBankAccountPaymentMethodOptions {
	if o == nil || isNil(o.UsBankAccount) {
		var ret CheckoutUsBankAccountPaymentMethodOptions
		return ret
	}
	return *o.UsBankAccount
}

// GetUsBankAccountOk returns a tuple with the UsBankAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CheckoutSessionPaymentMethodOptions) GetUsBankAccountOk() (*CheckoutUsBankAccountPaymentMethodOptions, bool) {
	if o == nil || isNil(o.UsBankAccount) {
		return nil, false
	}
	return o.UsBankAccount, true
}

// HasUsBankAccount returns a boolean if a field has been set.
func (o *CheckoutSessionPaymentMethodOptions) HasUsBankAccount() bool {
	if o != nil && !isNil(o.UsBankAccount) {
		return true
	}

	return false
}

// SetUsBankAccount gets a reference to the given CheckoutUsBankAccountPaymentMethodOptions and assigns it to the UsBankAccount field.
func (o *CheckoutSessionPaymentMethodOptions) SetUsBankAccount(v CheckoutUsBankAccountPaymentMethodOptions) {
	o.UsBankAccount = &v
}

func (o CheckoutSessionPaymentMethodOptions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.AcssDebit) {
		toSerialize["acss_debit"] = o.AcssDebit
	}
	if !isNil(o.Affirm) {
		toSerialize["affirm"] = o.Affirm
	}
	if !isNil(o.AfterpayClearpay) {
		toSerialize["afterpay_clearpay"] = o.AfterpayClearpay
	}
	if !isNil(o.Alipay) {
		toSerialize["alipay"] = o.Alipay
	}
	if !isNil(o.AuBecsDebit) {
		toSerialize["au_becs_debit"] = o.AuBecsDebit
	}
	if !isNil(o.BacsDebit) {
		toSerialize["bacs_debit"] = o.BacsDebit
	}
	if !isNil(o.Bancontact) {
		toSerialize["bancontact"] = o.Bancontact
	}
	if !isNil(o.Boleto) {
		toSerialize["boleto"] = o.Boleto
	}
	if !isNil(o.Card) {
		toSerialize["card"] = o.Card
	}
	if !isNil(o.CustomerBalance) {
		toSerialize["customer_balance"] = o.CustomerBalance
	}
	if !isNil(o.Eps) {
		toSerialize["eps"] = o.Eps
	}
	if !isNil(o.Fpx) {
		toSerialize["fpx"] = o.Fpx
	}
	if !isNil(o.Giropay) {
		toSerialize["giropay"] = o.Giropay
	}
	if !isNil(o.Grabpay) {
		toSerialize["grabpay"] = o.Grabpay
	}
	if !isNil(o.Ideal) {
		toSerialize["ideal"] = o.Ideal
	}
	if !isNil(o.Klarna) {
		toSerialize["klarna"] = o.Klarna
	}
	if !isNil(o.Konbini) {
		toSerialize["konbini"] = o.Konbini
	}
	if !isNil(o.Oxxo) {
		toSerialize["oxxo"] = o.Oxxo
	}
	if !isNil(o.P24) {
		toSerialize["p24"] = o.P24
	}
	if !isNil(o.Paynow) {
		toSerialize["paynow"] = o.Paynow
	}
	if !isNil(o.Pix) {
		toSerialize["pix"] = o.Pix
	}
	if !isNil(o.SepaDebit) {
		toSerialize["sepa_debit"] = o.SepaDebit
	}
	if !isNil(o.Sofort) {
		toSerialize["sofort"] = o.Sofort
	}
	if !isNil(o.UsBankAccount) {
		toSerialize["us_bank_account"] = o.UsBankAccount
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutSessionPaymentMethodOptions struct {
	value *CheckoutSessionPaymentMethodOptions
	isSet bool
}

func (v NullableCheckoutSessionPaymentMethodOptions) Get() *CheckoutSessionPaymentMethodOptions {
	return v.value
}

func (v *NullableCheckoutSessionPaymentMethodOptions) Set(val *CheckoutSessionPaymentMethodOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionPaymentMethodOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionPaymentMethodOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionPaymentMethodOptions(val *CheckoutSessionPaymentMethodOptions) *NullableCheckoutSessionPaymentMethodOptions {
	return &NullableCheckoutSessionPaymentMethodOptions{value: val, isSet: true}
}

func (v NullableCheckoutSessionPaymentMethodOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionPaymentMethodOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
