package shipping

import "github.com/driver005/gateway/utils/tax"

// ShippingRateDeliveryEstimateBound
type ShippingRateDeliveryEstimateBound struct {
	// A unit of time.
	Unit string `json:"unit"`
	// Must be greater than 0.
	Value int32 `json:"value"`
}

// ShippingRateDeliveryEstimate
type ShippingRateDeliveryEstimate struct {
	Maximum *ShippingRateDeliveryEstimateBound `json:"maximum,omitempty"`
	Minimum *ShippingRateDeliveryEstimateBound `json:"minimum,omitempty"`
}

// ShippingRateCurrencyOption
type ShippingRateCurrencyOption struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int32 `json:"amount"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior string `json:"tax_behavior"`
}

// ShippingRateFixedAmount
type ShippingRateFixedAmount struct {
	// A non-negative integer in cents representing how much to charge.
	Amount int32 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *ShippingRateCurrencyOption `json:"currency_options,omitempty"`
}

// ShippingRate Shipping rates describe the price of shipping presented to your customers and can be applied to [Checkout Sessions](https://stripe.com/docs/payments/checkout/shipping) and [Orders](https://stripe.com/docs/orders/shipping) to collect shipping costs.
type ShippingRate struct {
	// Whether the shipping rate can be used for new purchases. Defaults to `true`.
	Active           bool                          `json:"active"`
	DeliveryEstimate *ShippingRateDeliveryEstimate `json:"delivery_estimate,omitempty"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName string                   `json:"display_name,omitempty"`
	FixedAmount *ShippingRateFixedAmount `json:"fixed_amount,omitempty"`
	TaxBehavior string                   `json:"tax_behavior,omitempty"`
	TaxCode     *tax.TaxCode             `json:"tax_code,omitempty"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type string `json:"type"`
}
