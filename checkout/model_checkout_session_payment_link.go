package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionPaymentLink The ID of the Payment Link that created this Session.
type CheckoutSessionPaymentLink struct {
	PaymentLink *PaymentLink
	string      *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionPaymentLink) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentLink
	err = json.Unmarshal(data, &dst.PaymentLink)
	if err == nil {
		jsonPaymentLink, _ := json.Marshal(dst.PaymentLink)
		if string(jsonPaymentLink) == "{}" { // empty struct
			dst.PaymentLink = nil
		} else {
			return nil // data stored in dst.PaymentLink, return on the first match
		}
	} else {
		dst.PaymentLink = nil
	}

	// try to unmarshal JSON data into string
	err = json.Unmarshal(data, &dst.string)
	if err == nil {
		jsonstring, _ := json.Marshal(dst.string)
		if string(jsonstring) == "{}" { // empty struct
			dst.string = nil
		} else {
			return nil // data stored in dst.string, return on the first match
		}
	} else {
		dst.string = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionPaymentLink)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionPaymentLink) MarshalJSON() ([]byte, error) {
	if src.PaymentLink != nil {
		return json.Marshal(&src.PaymentLink)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionPaymentLink struct {
	value *CheckoutSessionPaymentLink
	isSet bool
}

func (v NullableCheckoutSessionPaymentLink) Get() *CheckoutSessionPaymentLink {
	return v.value
}

func (v *NullableCheckoutSessionPaymentLink) Set(val *CheckoutSessionPaymentLink) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionPaymentLink) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionPaymentLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionPaymentLink(val *CheckoutSessionPaymentLink) *NullableCheckoutSessionPaymentLink {
	return &NullableCheckoutSessionPaymentLink{value: val, isSet: true}
}

func (v NullableCheckoutSessionPaymentLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionPaymentLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
