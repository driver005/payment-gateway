package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionShippingDetails Shipping information for this Checkout Session.
type CheckoutSessionShippingDetails struct {
	Shipping *Shipping
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionShippingDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Shipping
	err = json.Unmarshal(data, &dst.Shipping)
	if err == nil {
		jsonShipping, _ := json.Marshal(dst.Shipping)
		if string(jsonShipping) == "{}" { // empty struct
			dst.Shipping = nil
		} else {
			return nil // data stored in dst.Shipping, return on the first match
		}
	} else {
		dst.Shipping = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionShippingDetails)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionShippingDetails) MarshalJSON() ([]byte, error) {
	if src.Shipping != nil {
		return json.Marshal(&src.Shipping)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionShippingDetails struct {
	value *CheckoutSessionShippingDetails
	isSet bool
}

func (v NullableCheckoutSessionShippingDetails) Get() *CheckoutSessionShippingDetails {
	return v.value
}

func (v *NullableCheckoutSessionShippingDetails) Set(val *CheckoutSessionShippingDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionShippingDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionShippingDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionShippingDetails(val *CheckoutSessionShippingDetails) *NullableCheckoutSessionShippingDetails {
	return &NullableCheckoutSessionShippingDetails{value: val, isSet: true}
}

func (v NullableCheckoutSessionShippingDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionShippingDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
