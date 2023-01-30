package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionConsent Results of `consent_collection` for this session.
type CheckoutSessionConsent struct {
	PaymentPagesCheckoutSessionConsent *PaymentPagesCheckoutSessionConsent
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionConsent) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionConsent
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionConsent)
	if err == nil {
		jsonPaymentPagesCheckoutSessionConsent, _ := json.Marshal(dst.PaymentPagesCheckoutSessionConsent)
		if string(jsonPaymentPagesCheckoutSessionConsent) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionConsent = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionConsent, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionConsent = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionConsent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionConsent) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionConsent != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionConsent)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionConsent struct {
	value *CheckoutSessionConsent
	isSet bool
}

func (v NullableCheckoutSessionConsent) Get() *CheckoutSessionConsent {
	return v.value
}

func (v *NullableCheckoutSessionConsent) Set(val *CheckoutSessionConsent) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionConsent) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionConsent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionConsent(val *CheckoutSessionConsent) *NullableCheckoutSessionConsent {
	return &NullableCheckoutSessionConsent{value: val, isSet: true}
}

func (v NullableCheckoutSessionConsent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionConsent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
