package models

type Customer struct {
	ID              string      `json:"id"`
	Object          string      `json:"object"`
	Address         interface{} `json:"address"`
	Balance         int         `json:"balance"`
	Created         int         `json:"created"`
	Currency        string      `json:"currency"`
	DefaultSource   string      `json:"default_source"`
	Delinquent      bool        `json:"delinquent"`
	Description     string      `json:"description"`
	Discount        interface{} `json:"discount"`
	Email           interface{} `json:"email"`
	InvoicePrefix   string      `json:"invoice_prefix"`
	InvoiceSettings struct {
		CustomFields         interface{} `json:"custom_fields"`
		DefaultPaymentMethod interface{} `json:"default_payment_method"`
		Footer               interface{} `json:"footer"`
	} `json:"invoice_settings"`
	Livemode bool `json:"livemode"`
	Metadata struct {
		OrderID string `json:"order_id"`
	} `json:"metadata"`
	Name                interface{}   `json:"name"`
	NextInvoiceSequence int           `json:"next_invoice_sequence"`
	Phone               interface{}   `json:"phone"`
	PreferredLocales    []interface{} `json:"preferred_locales"`
	Shipping            interface{}   `json:"shipping"`
	TaxExempt           string        `json:"tax_exempt"`
	TestClock           interface{}   `json:"test_clock"`
}