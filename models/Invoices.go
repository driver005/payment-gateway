package models

type Invoices struct {
	ID                   string      `json:"id"`
	Object               string      `json:"object"`
	AccountCountry       string      `json:"account_country"`
	AccountName          string      `json:"account_name"`
	AccountTaxIds        interface{} `json:"account_tax_ids"`
	AmountDue            int         `json:"amount_due"`
	AmountPaid           int         `json:"amount_paid"`
	AmountRemaining      int         `json:"amount_remaining"`
	ApplicationFeeAmount interface{} `json:"application_fee_amount"`
	AttemptCount         int         `json:"attempt_count"`
	Attempted            bool        `json:"attempted"`
	AutoAdvance          bool        `json:"auto_advance"`
	AutomaticTax         struct {
		Enabled bool        `json:"enabled"`
		Status  interface{} `json:"status"`
	} `json:"automatic_tax"`
	BillingReason         string        `json:"billing_reason"`
	Charge                interface{}   `json:"charge"`
	CollectionMethod      string        `json:"collection_method"`
	Created               int           `json:"created"`
	Currency              string        `json:"currency"`
	CustomFields          interface{}   `json:"custom_fields"`
	Customer              string        `json:"customer"`
	CustomerAddress       interface{}   `json:"customer_address"`
	CustomerEmail         interface{}   `json:"customer_email"`
	CustomerName          interface{}   `json:"customer_name"`
	CustomerPhone         interface{}   `json:"customer_phone"`
	CustomerShipping      interface{}   `json:"customer_shipping"`
	CustomerTaxExempt     string        `json:"customer_tax_exempt"`
	CustomerTaxIds        []interface{} `json:"customer_tax_ids"`
	DefaultPaymentMethod  interface{}   `json:"default_payment_method"`
	DefaultSource         interface{}   `json:"default_source"`
	DefaultTaxRates       []interface{} `json:"default_tax_rates"`
	Description           interface{}   `json:"description"`
	Discount              interface{}   `json:"discount"`
	Discounts             []interface{} `json:"discounts"`
	DueDate               interface{}   `json:"due_date"`
	EndingBalance         interface{}   `json:"ending_balance"`
	Footer                interface{}   `json:"footer"`
	HostedInvoiceURL      interface{}   `json:"hosted_invoice_url"`
	InvoicePdf            interface{}   `json:"invoice_pdf"`
	LastFinalizationError interface{}   `json:"last_finalization_error"`
	Lines                 struct {
		Object string `json:"object"`
		Data   []struct {
			ID              string        `json:"id"`
			Object          string        `json:"object"`
			Amount          int           `json:"amount"`
			Currency        string        `json:"currency"`
			Description     string        `json:"description"`
			DiscountAmounts []interface{} `json:"discount_amounts"`
			Discountable    bool          `json:"discountable"`
			Discounts       []interface{} `json:"discounts"`
			InvoiceItem     string        `json:"invoice_item"`
			Livemode        bool          `json:"livemode"`
			Metadata        struct {
			} `json:"metadata"`
			Period struct {
				End   int `json:"end"`
				Start int `json:"start"`
			} `json:"period"`
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
				} `json:"metadata"`
				Nickname          interface{} `json:"nickname"`
				Product           string      `json:"product"`
				Recurring         interface{} `json:"recurring"`
				TaxBehavior       string      `json:"tax_behavior"`
				TiersMode         interface{} `json:"tiers_mode"`
				TransformQuantity interface{} `json:"transform_quantity"`
				Type              string      `json:"type"`
				UnitAmount        int         `json:"unit_amount"`
				UnitAmountDecimal string      `json:"unit_amount_decimal"`
			} `json:"price"`
			Proration        bool `json:"proration"`
			ProrationDetails struct {
				CreditedItems interface{} `json:"credited_items"`
			} `json:"proration_details"`
			Quantity     int           `json:"quantity"`
			Subscription interface{}   `json:"subscription"`
			TaxAmounts   []interface{} `json:"tax_amounts"`
			TaxRates     []interface{} `json:"tax_rates"`
			Type         string        `json:"type"`
		} `json:"data"`
		HasMore bool   `json:"has_more"`
		URL     string `json:"url"`
	} `json:"lines"`
	Livemode bool `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	NextPaymentAttempt int         `json:"next_payment_attempt"`
	Number             string      `json:"number"`
	OnBehalfOf         interface{} `json:"on_behalf_of"`
	Paid               bool        `json:"paid"`
	PaidOutOfBand      bool        `json:"paid_out_of_band"`
	PaymentIntent      interface{} `json:"payment_intent"`
	PaymentSettings    struct {
		PaymentMethodOptions interface{} `json:"payment_method_options"`
		PaymentMethodTypes   interface{} `json:"payment_method_types"`
	} `json:"payment_settings"`
	PeriodEnd                    int         `json:"period_end"`
	PeriodStart                  int         `json:"period_start"`
	PostPaymentCreditNotesAmount int         `json:"post_payment_credit_notes_amount"`
	PrePaymentCreditNotesAmount  int         `json:"pre_payment_credit_notes_amount"`
	Quote                        interface{} `json:"quote"`
	ReceiptNumber                interface{} `json:"receipt_number"`
	StartingBalance              int         `json:"starting_balance"`
	StatementDescriptor          interface{} `json:"statement_descriptor"`
	Status                       string      `json:"status"`
	StatusTransitions            struct {
		FinalizedAt           interface{} `json:"finalized_at"`
		MarkedUncollectibleAt interface{} `json:"marked_uncollectible_at"`
		PaidAt                interface{} `json:"paid_at"`
		VoidedAt              interface{} `json:"voided_at"`
	} `json:"status_transitions"`
	Subscription         interface{}   `json:"subscription"`
	Subtotal             int           `json:"subtotal"`
	Tax                  interface{}   `json:"tax"`
	TestClock            interface{}   `json:"test_clock"`
	Total                int           `json:"total"`
	TotalDiscountAmounts []interface{} `json:"total_discount_amounts"`
	TotalTaxAmounts      []interface{} `json:"total_tax_amounts"`
	TransferData         interface{}   `json:"transfer_data"`
	WebhooksDeliveredAt  interface{}   `json:"webhooks_delivered_at"`
}