package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionTotalDetails Tax and discount details for the computed total amount.
type CheckoutSessionTotalDetails struct {
	PaymentPagesCheckoutSessionTotalDetails *PaymentPagesCheckoutSessionTotalDetails
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionTotalDetails) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionTotalDetails
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionTotalDetails)
	if err == nil {
		jsonPaymentPagesCheckoutSessionTotalDetails, _ := json.Marshal(dst.PaymentPagesCheckoutSessionTotalDetails)
		if string(jsonPaymentPagesCheckoutSessionTotalDetails) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionTotalDetails = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionTotalDetails, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionTotalDetails = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionTotalDetails)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionTotalDetails) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionTotalDetails != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionTotalDetails)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionTotalDetails struct {
	value *CheckoutSessionTotalDetails
	isSet bool
}

func (v NullableCheckoutSessionTotalDetails) Get() *CheckoutSessionTotalDetails {
	return v.value
}

func (v *NullableCheckoutSessionTotalDetails) Set(val *CheckoutSessionTotalDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionTotalDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionTotalDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionTotalDetails(val *CheckoutSessionTotalDetails) *NullableCheckoutSessionTotalDetails {
	return &NullableCheckoutSessionTotalDetails{value: val, isSet: true}
}

func (v NullableCheckoutSessionTotalDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionTotalDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
