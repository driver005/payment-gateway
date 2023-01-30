package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionSubscription The ID of the subscription for Checkout Sessions in `subscription` mode.
type CheckoutSessionSubscription struct {
	Subscription *Subscription
	string       *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionSubscription) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Subscription
	err = json.Unmarshal(data, &dst.Subscription)
	if err == nil {
		jsonSubscription, _ := json.Marshal(dst.Subscription)
		if string(jsonSubscription) == "{}" { // empty struct
			dst.Subscription = nil
		} else {
			return nil // data stored in dst.Subscription, return on the first match
		}
	} else {
		dst.Subscription = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionSubscription)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionSubscription) MarshalJSON() ([]byte, error) {
	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionSubscription struct {
	value *CheckoutSessionSubscription
	isSet bool
}

func (v NullableCheckoutSessionSubscription) Get() *CheckoutSessionSubscription {
	return v.value
}

func (v *NullableCheckoutSessionSubscription) Set(val *CheckoutSessionSubscription) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionSubscription) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionSubscription) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionSubscription(val *CheckoutSessionSubscription) *NullableCheckoutSessionSubscription {
	return &NullableCheckoutSessionSubscription{value: val, isSet: true}
}

func (v NullableCheckoutSessionSubscription) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionSubscription) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
