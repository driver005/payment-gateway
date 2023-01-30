package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionPaymentIntent The ID of the PaymentIntent for Checkout Sessions in `payment` mode.
type CheckoutSessionPaymentIntent struct {
	PaymentIntent *PaymentIntent
	string        *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionPaymentIntent) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentIntent
	err = json.Unmarshal(data, &dst.PaymentIntent)
	if err == nil {
		jsonPaymentIntent, _ := json.Marshal(dst.PaymentIntent)
		if string(jsonPaymentIntent) == "{}" { // empty struct
			dst.PaymentIntent = nil
		} else {
			return nil // data stored in dst.PaymentIntent, return on the first match
		}
	} else {
		dst.PaymentIntent = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionPaymentIntent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionPaymentIntent) MarshalJSON() ([]byte, error) {
	if src.PaymentIntent != nil {
		return json.Marshal(&src.PaymentIntent)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionPaymentIntent struct {
	value *CheckoutSessionPaymentIntent
	isSet bool
}

func (v NullableCheckoutSessionPaymentIntent) Get() *CheckoutSessionPaymentIntent {
	return v.value
}

func (v *NullableCheckoutSessionPaymentIntent) Set(val *CheckoutSessionPaymentIntent) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionPaymentIntent) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionPaymentIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionPaymentIntent(val *CheckoutSessionPaymentIntent) *NullableCheckoutSessionPaymentIntent {
	return &NullableCheckoutSessionPaymentIntent{value: val, isSet: true}
}

func (v NullableCheckoutSessionPaymentIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionPaymentIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
