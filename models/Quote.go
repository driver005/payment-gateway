package models

type Quote struct {
	ID                    string      `json:"id"`
	Object                string      `json:"object"`
	AmountSubtotal        int         `json:"amount_subtotal"`
	AmountTotal           int         `json:"amount_total"`
	ApplicationFeeAmount  interface{} `json:"application_fee_amount"`
	ApplicationFeePercent interface{} `json:"application_fee_percent"`
	AutomaticTax          struct {
		Enabled bool        `json:"enabled"`
		Status  interface{} `json:"status"`
	} `json:"automatic_tax"`
	CollectionMethod string `json:"collection_method"`
	Computed         struct {
		Recurring interface{} `json:"recurring"`
		Upfront   struct {
			AmountSubtotal int `json:"amount_subtotal"`
			AmountTotal    int `json:"amount_total"`
			TotalDetails   struct {
				AmountDiscount int `json:"amount_discount"`
				AmountShipping int `json:"amount_shipping"`
				AmountTax      int `json:"amount_tax"`
			} `json:"total_details"`
		} `json:"upfront"`
	} `json:"computed"`
	Created         int           `json:"created"`
	Currency        string        `json:"currency"`
	Customer        string        `json:"customer"`
	DefaultTaxRates []interface{} `json:"default_tax_rates"`
	Description     interface{}   `json:"description"`
	Discounts       []interface{} `json:"discounts"`
	ExpiresAt       int           `json:"expires_at"`
	Footer          interface{}   `json:"footer"`
	FromQuote       interface{}   `json:"from_quote"`
	Header          interface{}   `json:"header"`
	Invoice         interface{}   `json:"invoice"`
	InvoiceSettings struct {
		DaysUntilDue interface{} `json:"days_until_due"`
	} `json:"invoice_settings"`
	Livemode bool `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	Number            interface{} `json:"number"`
	OnBehalfOf        interface{} `json:"on_behalf_of"`
	Status            string      `json:"status"`
	StatusTransitions struct {
		AcceptedAt  interface{} `json:"accepted_at"`
		CanceledAt  interface{} `json:"canceled_at"`
		FinalizedAt interface{} `json:"finalized_at"`
	} `json:"status_transitions"`
	Subscription     interface{} `json:"subscription"`
	SubscriptionData struct {
		EffectiveDate   interface{} `json:"effective_date"`
		TrialPeriodDays interface{} `json:"trial_period_days"`
	} `json:"subscription_data"`
	SubscriptionSchedule interface{} `json:"subscription_schedule"`
	TestClock            interface{} `json:"test_clock"`
	TotalDetails         struct {
		AmountDiscount int `json:"amount_discount"`
		AmountShipping int `json:"amount_shipping"`
		AmountTax      int `json:"amount_tax"`
	} `json:"total_details"`
	TransferData interface{} `json:"transfer_data"`
}