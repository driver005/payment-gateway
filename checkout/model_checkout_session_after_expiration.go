package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionAfterExpiration When set, provides configuration for actions to take if this Checkout Session expires.
type CheckoutSessionAfterExpiration struct {
	PaymentPagesCheckoutSessionAfterExpiration *PaymentPagesCheckoutSessionAfterExpiration
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionAfterExpiration) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionAfterExpiration
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionAfterExpiration)
	if err == nil {
		jsonPaymentPagesCheckoutSessionAfterExpiration, _ := json.Marshal(dst.PaymentPagesCheckoutSessionAfterExpiration)
		if string(jsonPaymentPagesCheckoutSessionAfterExpiration) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionAfterExpiration = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionAfterExpiration, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionAfterExpiration = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionAfterExpiration)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionAfterExpiration) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionAfterExpiration != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionAfterExpiration)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionAfterExpiration struct {
	value *CheckoutSessionAfterExpiration
	isSet bool
}

func (v NullableCheckoutSessionAfterExpiration) Get() *CheckoutSessionAfterExpiration {
	return v.value
}

func (v *NullableCheckoutSessionAfterExpiration) Set(val *CheckoutSessionAfterExpiration) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionAfterExpiration) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionAfterExpiration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionAfterExpiration(val *CheckoutSessionAfterExpiration) *NullableCheckoutSessionAfterExpiration {
	return &NullableCheckoutSessionAfterExpiration{value: val, isSet: true}
}

func (v NullableCheckoutSessionAfterExpiration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionAfterExpiration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
