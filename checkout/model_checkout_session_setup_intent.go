package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionSetupIntent The ID of the SetupIntent for Checkout Sessions in `setup` mode.
type CheckoutSessionSetupIntent struct {
	SetupIntent *SetupIntent
	string      *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionSetupIntent) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into SetupIntent
	err = json.Unmarshal(data, &dst.SetupIntent)
	if err == nil {
		jsonSetupIntent, _ := json.Marshal(dst.SetupIntent)
		if string(jsonSetupIntent) == "{}" { // empty struct
			dst.SetupIntent = nil
		} else {
			return nil // data stored in dst.SetupIntent, return on the first match
		}
	} else {
		dst.SetupIntent = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionSetupIntent)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionSetupIntent) MarshalJSON() ([]byte, error) {
	if src.SetupIntent != nil {
		return json.Marshal(&src.SetupIntent)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionSetupIntent struct {
	value *CheckoutSessionSetupIntent
	isSet bool
}

func (v NullableCheckoutSessionSetupIntent) Get() *CheckoutSessionSetupIntent {
	return v.value
}

func (v *NullableCheckoutSessionSetupIntent) Set(val *CheckoutSessionSetupIntent) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionSetupIntent) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionSetupIntent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionSetupIntent(val *CheckoutSessionSetupIntent) *NullableCheckoutSessionSetupIntent {
	return &NullableCheckoutSessionSetupIntent{value: val, isSet: true}
}

func (v NullableCheckoutSessionSetupIntent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionSetupIntent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
