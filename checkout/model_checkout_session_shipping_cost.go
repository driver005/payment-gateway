package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionShippingCost The details of the customer cost of shipping, including the customer chosen ShippingRate.
type CheckoutSessionShippingCost struct {
	PaymentPagesCheckoutSessionShippingCost *PaymentPagesCheckoutSessionShippingCost
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionShippingCost) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionShippingCost
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionShippingCost)
	if err == nil {
		jsonPaymentPagesCheckoutSessionShippingCost, _ := json.Marshal(dst.PaymentPagesCheckoutSessionShippingCost)
		if string(jsonPaymentPagesCheckoutSessionShippingCost) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionShippingCost = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionShippingCost, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionShippingCost = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionShippingCost)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionShippingCost) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionShippingCost != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionShippingCost)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionShippingCost struct {
	value *CheckoutSessionShippingCost
	isSet bool
}

func (v NullableCheckoutSessionShippingCost) Get() *CheckoutSessionShippingCost {
	return v.value
}

func (v *NullableCheckoutSessionShippingCost) Set(val *CheckoutSessionShippingCost) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionShippingCost) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionShippingCost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionShippingCost(val *CheckoutSessionShippingCost) *NullableCheckoutSessionShippingCost {
	return &NullableCheckoutSessionShippingCost{value: val, isSet: true}
}

func (v NullableCheckoutSessionShippingCost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionShippingCost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
