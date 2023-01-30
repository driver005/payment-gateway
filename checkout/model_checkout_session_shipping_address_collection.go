package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionShippingAddressCollection When set, provides configuration for Checkout to collect a shipping address from a customer.
type CheckoutSessionShippingAddressCollection struct {
	PaymentPagesCheckoutSessionShippingAddressCollection *PaymentPagesCheckoutSessionShippingAddressCollection
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionShippingAddressCollection) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionShippingAddressCollection
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionShippingAddressCollection)
	if err == nil {
		jsonPaymentPagesCheckoutSessionShippingAddressCollection, _ := json.Marshal(dst.PaymentPagesCheckoutSessionShippingAddressCollection)
		if string(jsonPaymentPagesCheckoutSessionShippingAddressCollection) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionShippingAddressCollection = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionShippingAddressCollection, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionShippingAddressCollection = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionShippingAddressCollection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionShippingAddressCollection) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionShippingAddressCollection != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionShippingAddressCollection)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionShippingAddressCollection struct {
	value *CheckoutSessionShippingAddressCollection
	isSet bool
}

func (v NullableCheckoutSessionShippingAddressCollection) Get() *CheckoutSessionShippingAddressCollection {
	return v.value
}

func (v *NullableCheckoutSessionShippingAddressCollection) Set(val *CheckoutSessionShippingAddressCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionShippingAddressCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionShippingAddressCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionShippingAddressCollection(val *CheckoutSessionShippingAddressCollection) *NullableCheckoutSessionShippingAddressCollection {
	return &NullableCheckoutSessionShippingAddressCollection{value: val, isSet: true}
}

func (v NullableCheckoutSessionShippingAddressCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionShippingAddressCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
