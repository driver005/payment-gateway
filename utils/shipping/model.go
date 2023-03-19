package shipping

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/tax"
)

// ShippingRateDeliveryEstimateBound
type ShippingRateDeliveryEstimateBound struct {
	core.Model

	// A unit of time.
	Unit string `json:"unit,omitempty"`
	// Must be greater than 0.
	Value int `json:"value,omitempty"`
}

// ShippingRateDeliveryEstimate
type ShippingRateDeliveryEstimate struct {
	core.Model

	Maximum *ShippingRateDeliveryEstimateBound `json:"maximum,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	Minimum *ShippingRateDeliveryEstimateBound `json:"minimum,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

// ShippingRateCurrencyOption
type ShippingRateCurrencyOption struct {
	core.Model

	// A non-negative integer in cents representing how much to charge.
	Amount int `json:"amount,omitempty"`
	// Specifies whether the rate is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`.
	TaxBehavior TaxBehavior `json:"tax_behavior,omitempty"`
}

// ShippingRateFixedAmount
type ShippingRateFixedAmount struct {
	core.Model

	// A non-negative integer in cents representing how much to charge.
	Amount int `json:"amount,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency string `json:"currency,omitempty"`
	// Shipping rates defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *ShippingRateCurrencyOption `json:"currency_options,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
}

// ShippingRate Shipping rates describe the price of shipping presented to your customers and can be applied to [Checkout Sessions](https://stripe.com/docs/payments/checkout/shipping) and [Orders](https://stripe.com/docs/orders/shipping) to collect shipping costs.
type ShippingRate struct {
	core.Model

	// Whether the shipping rate can be used for new purchases. Defaults to `true`.
	Active           bool                          `json:"active"`
	DeliveryEstimate *ShippingRateDeliveryEstimate `json:"delivery_estimate,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The name of the shipping rate, meant to be displayable to the customer. This will appear on CheckoutSessions.
	DisplayName string                   `json:"display_name,omitempty"`
	FixedAmount *ShippingRateFixedAmount `json:"fixed_amount,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	TaxBehavior TaxBehavior              `json:"tax_behavior,omitempty"`
	TaxCode     *tax.TaxCode             `json:"tax_code,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	// The type of calculation to use on the shipping rate. Can only be `fixed_amount` for now.
	Type Type `json:"type"`
}
