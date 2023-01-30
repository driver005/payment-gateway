package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionCustomer The ID of the customer for this Session. For Checkout Sessions in `payment` or `subscription` mode, Checkout will create a new customer object based on information provided during the payment flow unless an existing customer was provided when the Session was created.
type CheckoutSessionCustomer struct {
	Customer        *Customer
	DeletedCustomer *DeletedCustomer
	string          *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionCustomer) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Customer
	err = json.Unmarshal(data, &dst.Customer)
	if err == nil {
		jsonCustomer, _ := json.Marshal(dst.Customer)
		if string(jsonCustomer) == "{}" { // empty struct
			dst.Customer = nil
		} else {
			return nil // data stored in dst.Customer, return on the first match
		}
	} else {
		dst.Customer = nil
	}

	// try to unmarshal JSON data into DeletedCustomer
	err = json.Unmarshal(data, &dst.DeletedCustomer)
	if err == nil {
		jsonDeletedCustomer, _ := json.Marshal(dst.DeletedCustomer)
		if string(jsonDeletedCustomer) == "{}" { // empty struct
			dst.DeletedCustomer = nil
		} else {
			return nil // data stored in dst.DeletedCustomer, return on the first match
		}
	} else {
		dst.DeletedCustomer = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionCustomer)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionCustomer) MarshalJSON() ([]byte, error) {
	if src.Customer != nil {
		return json.Marshal(&src.Customer)
	}

	if src.DeletedCustomer != nil {
		return json.Marshal(&src.DeletedCustomer)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionCustomer struct {
	value *CheckoutSessionCustomer
	isSet bool
}

func (v NullableCheckoutSessionCustomer) Get() *CheckoutSessionCustomer {
	return v.value
}

func (v *NullableCheckoutSessionCustomer) Set(val *CheckoutSessionCustomer) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionCustomer) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionCustomer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionCustomer(val *CheckoutSessionCustomer) *NullableCheckoutSessionCustomer {
	return &NullableCheckoutSessionCustomer{value: val, isSet: true}
}

func (v NullableCheckoutSessionCustomer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionCustomer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
