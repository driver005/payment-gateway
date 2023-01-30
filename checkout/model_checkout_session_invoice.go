package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionInvoice ID of the invoice created by the Checkout Session, if it exists.
type CheckoutSessionInvoice struct {
	Invoice *Invoice
	string  *string
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionInvoice) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into Invoice
	err = json.Unmarshal(data, &dst.Invoice)
	if err == nil {
		jsonInvoice, _ := json.Marshal(dst.Invoice)
		if string(jsonInvoice) == "{}" { // empty struct
			dst.Invoice = nil
		} else {
			return nil // data stored in dst.Invoice, return on the first match
		}
	} else {
		dst.Invoice = nil
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

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionInvoice)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionInvoice) MarshalJSON() ([]byte, error) {
	if src.Invoice != nil {
		return json.Marshal(&src.Invoice)
	}

	if src.string != nil {
		return json.Marshal(&src.string)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionInvoice struct {
	value *CheckoutSessionInvoice
	isSet bool
}

func (v NullableCheckoutSessionInvoice) Get() *CheckoutSessionInvoice {
	return v.value
}

func (v *NullableCheckoutSessionInvoice) Set(val *CheckoutSessionInvoice) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionInvoice) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionInvoice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionInvoice(val *CheckoutSessionInvoice) *NullableCheckoutSessionInvoice {
	return &NullableCheckoutSessionInvoice{value: val, isSet: true}
}

func (v NullableCheckoutSessionInvoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionInvoice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
