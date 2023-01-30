package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionCustomerDetails The customer details including the customer's tax exempt status and the customer's tax IDs. Only the customer's email is present on Sessions in `setup` mode.
type CheckoutSessionCustomerDetails struct {
	PaymentPagesCheckoutSessionCustomerDetails *PaymentPagesCheckoutSessionCustomerDetails
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionCustomerDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionCustomerDetails
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionCustomerDetails)
	if err == nil {
		jsonPaymentPagesCheckoutSessionCustomerDetails, _ := json.Marshal(dst.PaymentPagesCheckoutSessionCustomerDetails)
		if string(jsonPaymentPagesCheckoutSessionCustomerDetails) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionCustomerDetails = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionCustomerDetails, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionCustomerDetails = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionCustomerDetails)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionCustomerDetails) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionCustomerDetails != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionCustomerDetails)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionCustomerDetails struct {
	value *CheckoutSessionCustomerDetails
	isSet bool
}

func (v NullableCheckoutSessionCustomerDetails) Get() *CheckoutSessionCustomerDetails {
	return v.value
}

func (v *NullableCheckoutSessionCustomerDetails) Set(val *CheckoutSessionCustomerDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionCustomerDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionCustomerDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionCustomerDetails(val *CheckoutSessionCustomerDetails) *NullableCheckoutSessionCustomerDetails {
	return &NullableCheckoutSessionCustomerDetails{value: val, isSet: true}
}

func (v NullableCheckoutSessionCustomerDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionCustomerDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
