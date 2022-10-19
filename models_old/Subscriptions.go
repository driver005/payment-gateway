package models

type Subscription struct {
	ID                    string      `json:"id"`
	Object                string      `json:"object"`
	ApplicationFeePercent interface{} `json:"application_fee_percent"`
	AutomaticTax          struct {
		Enabled bool `json:"enabled"`
	} `json:"automatic_tax"`
	BillingCycleAnchor   int           `json:"billing_cycle_anchor"`
	BillingThresholds    interface{}   `json:"billing_thresholds"`
	CancelAt             interface{}   `json:"cancel_at"`
	CancelAtPeriodEnd    bool          `json:"cancel_at_period_end"`
	CanceledAt           interface{}   `json:"canceled_at"`
	CollectionMethod     string        `json:"collection_method"`
	Created              int           `json:"created"`
	CurrentPeriodEnd     int           `json:"current_period_end"`
	CurrentPeriodStart   int           `json:"current_period_start"`
	Customer             string        `json:"customer"`
	DaysUntilDue         interface{}   `json:"days_until_due"`
	DefaultPaymentMethod string        `json:"default_payment_method"`
	DefaultSource        interface{}   `json:"default_source"`
	DefaultTaxRates      []interface{} `json:"default_tax_rates"`
	Discount             interface{}   `json:"discount"`
	EndedAt              interface{}   `json:"ended_at"`
	Items                struct {
		Object string `json:"object"`
		Data   []struct {
			ID                string      `json:"id"`
			Object            string      `json:"object"`
			BillingThresholds interface{} `json:"billing_thresholds"`
			Created           int         `json:"created"`
			Metadata          struct {
			} `json:"metadata"`
			Price struct {
				ID            string      `json:"id"`
				Object        string      `json:"object"`
				Active        bool        `json:"active"`
				BillingScheme string      `json:"billing_scheme"`
				Created       int         `json:"created"`
				Currency      string      `json:"currency"`
				Livemode      bool        `json:"livemode"`
				LookupKey     interface{} `json:"lookup_key"`
				Metadata      struct {
					Charset string `json:"charset"`
					Content string `json:"content"`
				} `json:"metadata"`
				Nickname  interface{} `json:"nickname"`
				Product   string      `json:"product"`
				Recurring struct {
					AggregateUsage interface{} `json:"aggregate_usage"`
					Interval       string      `json:"interval"`
					IntervalCount  int         `json:"interval_count"`
					UsageType      string      `json:"usage_type"`
				} `json:"recurring"`
				TaxBehavior       string      `json:"tax_behavior"`
				TiersMode         interface{} `json:"tiers_mode"`
				TransformQuantity interface{} `json:"transform_quantity"`
				Type              string      `json:"type"`
				UnitAmount        int         `json:"unit_amount"`
				UnitAmountDecimal string      `json:"unit_amount_decimal"`
			} `json:"price"`
			Quantity     int           `json:"quantity"`
			Subscription string        `json:"subscription"`
			TaxRates     []interface{} `json:"tax_rates"`
		} `json:"data"`
		HasMore bool   `json:"has_more"`
		URL     string `json:"url"`
	} `json:"items"`
	LatestInvoice string `json:"latest_invoice"`
	Livemode      bool   `json:"livemode"`
	Metadata      struct {
	} `json:"metadata"`
	NextPendingInvoiceItemInvoice interface{} `json:"next_pending_invoice_item_invoice"`
	PauseCollection               interface{} `json:"pause_collection"`
	PaymentSettings               struct {
		PaymentMethodOptions interface{} `json:"payment_method_options"`
		PaymentMethodTypes   interface{} `json:"payment_method_types"`
	} `json:"payment_settings"`
	PendingInvoiceItemInterval interface{} `json:"pending_invoice_item_interval"`
	PendingSetupIntent         interface{} `json:"pending_setup_intent"`
	PendingUpdate              interface{} `json:"pending_update"`
	Schedule                   interface{} `json:"schedule"`
	StartDate                  int         `json:"start_date"`
	Status                     string      `json:"status"`
	TestClock                  interface{} `json:"test_clock"`
	TransferData               interface{} `json:"transfer_data"`
	TrialEnd                   int         `json:"trial_end"`
	TrialStart                 int         `json:"trial_start"`
}
