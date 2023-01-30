package options

import "github.com/driver005/gateway/core"

// PaymentMethodOptionsCardMandateOptions
type PaymentMethodOptionsCardMandateOptions struct {
	core.Model

	// Amount to be charged for future payments.
	Amount int `json:"amount,omitempty"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType string `json:"amount_type,omitempty"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description string `json:"description,omitempty"`
	// End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.
	EndDate int `json:"end_date,omitempty"`
	// Specifies payment frequency. One of `day`, `week`, `month`, `year`, or `sporadic`.
	Interval string `json:"interval,omitempty"`
	// The number of intervals between payments. For example, `interval=month` and `interval_count=3` indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when `interval=sporadic`.
	IntervalCount int `json:"interval_count,omitempty"`
	// Unique identifier for the mandate or subscription.
	Reference string `json:"reference,omitempty"`
	// Start date of the mandate or subscription. Start date should not be lesser than yesterday.
	StartDate int `json:"start_date,omitempty"`
	// Specifies the type of mandates supported. Possible values are `india`.
	SupportedTypes []string `json:"supported_types,omitempty" database:"type:text[]"`
}
