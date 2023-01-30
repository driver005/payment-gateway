package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionInvoiceCreation Details on the state of invoice creation for the Checkout Session.
type CheckoutSessionInvoiceCreation struct {
	PaymentPagesCheckoutSessionInvoiceCreation *PaymentPagesCheckoutSessionInvoiceCreation
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionInvoiceCreation) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionInvoiceCreation
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionInvoiceCreation)
	if err == nil {
		jsonPaymentPagesCheckoutSessionInvoiceCreation, _ := json.Marshal(dst.PaymentPagesCheckoutSessionInvoiceCreation)
		if string(jsonPaymentPagesCheckoutSessionInvoiceCreation) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionInvoiceCreation = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionInvoiceCreation, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionInvoiceCreation = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionInvoiceCreation)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionInvoiceCreation) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionInvoiceCreation != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionInvoiceCreation)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionInvoiceCreation struct {
	value *CheckoutSessionInvoiceCreation
	isSet bool
}

func (v NullableCheckoutSessionInvoiceCreation) Get() *CheckoutSessionInvoiceCreation {
	return v.value
}

func (v *NullableCheckoutSessionInvoiceCreation) Set(val *CheckoutSessionInvoiceCreation) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionInvoiceCreation) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionInvoiceCreation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionInvoiceCreation(val *CheckoutSessionInvoiceCreation) *NullableCheckoutSessionInvoiceCreation {
	return &NullableCheckoutSessionInvoiceCreation{value: val, isSet: true}
}

func (v NullableCheckoutSessionInvoiceCreation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionInvoiceCreation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
