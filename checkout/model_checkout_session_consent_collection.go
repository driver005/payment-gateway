package checkout

import (
	"encoding/json"
	"fmt"
)

// CheckoutSessionConsentCollection When set, provides configuration for the Checkout Session to gather active consent from customers.
type CheckoutSessionConsentCollection struct {
	PaymentPagesCheckoutSessionConsentCollection *PaymentPagesCheckoutSessionConsentCollection
}

// Unmarshal JSON data into any of the pointers in the struct
func (dst *CheckoutSessionConsentCollection) UnmarshalJSON(data []byte) error {
	var err error
	// this object is nullable so check if the payload is null or empty string
	if string(data) == "" || string(data) == "{}" {
		return nil
	}

	// try to unmarshal JSON data into PaymentPagesCheckoutSessionConsentCollection
	err = json.Unmarshal(data, &dst.PaymentPagesCheckoutSessionConsentCollection)
	if err == nil {
		jsonPaymentPagesCheckoutSessionConsentCollection, _ := json.Marshal(dst.PaymentPagesCheckoutSessionConsentCollection)
		if string(jsonPaymentPagesCheckoutSessionConsentCollection) == "{}" { // empty struct
			dst.PaymentPagesCheckoutSessionConsentCollection = nil
		} else {
			return nil // data stored in dst.PaymentPagesCheckoutSessionConsentCollection, return on the first match
		}
	} else {
		dst.PaymentPagesCheckoutSessionConsentCollection = nil
	}

	return fmt.Errorf("data failed to match schemas in anyOf(CheckoutSessionConsentCollection)")
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src *CheckoutSessionConsentCollection) MarshalJSON() ([]byte, error) {
	if src.PaymentPagesCheckoutSessionConsentCollection != nil {
		return json.Marshal(&src.PaymentPagesCheckoutSessionConsentCollection)
	}

	return nil, nil // no data in anyOf schemas
}

type NullableCheckoutSessionConsentCollection struct {
	value *CheckoutSessionConsentCollection
	isSet bool
}

func (v NullableCheckoutSessionConsentCollection) Get() *CheckoutSessionConsentCollection {
	return v.value
}

func (v *NullableCheckoutSessionConsentCollection) Set(val *CheckoutSessionConsentCollection) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutSessionConsentCollection) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutSessionConsentCollection) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutSessionConsentCollection(val *CheckoutSessionConsentCollection) *NullableCheckoutSessionConsentCollection {
	return &NullableCheckoutSessionConsentCollection{value: val, isSet: true}
}

func (v NullableCheckoutSessionConsentCollection) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutSessionConsentCollection) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
