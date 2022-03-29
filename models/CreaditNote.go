package models

type CreditNote struct {
	ID                         string        `json:"id"`
	Object                     string        `json:"object"`
	Amount                     int           `json:"amount"`
	Created                    int           `json:"created"`
	Currency                   string        `json:"currency"`
	Customer                   string        `json:"customer"`
	CustomerBalanceTransaction interface{}   `json:"customer_balance_transaction"`
	DiscountAmount             int           `json:"discount_amount"`
	DiscountAmounts            []interface{} `json:"discount_amounts"`
	Invoice                    string        `json:"invoice"`
	Lines                      struct {
		Object string `json:"object"`
		Data   []struct {
			ID              string        `json:"id"`
			Object          string        `json:"object"`
			Amount          int           `json:"amount"`
			Description     string        `json:"description"`
			DiscountAmount  int           `json:"discount_amount"`
			DiscountAmounts []interface{} `json:"discount_amounts"`
			InvoiceLineItem string        `json:"invoice_line_item,omitempty"`
			Livemode        bool          `json:"livemode"`
			Quantity        int           `json:"quantity"`
			TaxAmounts      []struct {
				Amount    int    `json:"amount"`
				Inclusive bool   `json:"inclusive"`
				TaxRate   string `json:"tax_rate"`
			} `json:"tax_amounts"`
			TaxRates []struct {
				ID           string `json:"id"`
				Object       string `json:"object"`
				Active       bool   `json:"active"`
				Country      string `json:"country"`
				Created      int    `json:"created"`
				Description  string `json:"description"`
				DisplayName  string `json:"display_name"`
				Inclusive    bool   `json:"inclusive"`
				Jurisdiction string `json:"jurisdiction"`
				Livemode     bool   `json:"livemode"`
				Metadata     struct {
				} `json:"metadata"`
				Percentage float64     `json:"percentage"`
				State      interface{} `json:"state"`
				TaxType    string      `json:"tax_type"`
			} `json:"tax_rates"`
			Type              string      `json:"type"`
			UnitAmount        interface{} `json:"unit_amount"`
			UnitAmountDecimal interface{} `json:"unit_amount_decimal"`
		} `json:"data"`
		HasMore bool   `json:"has_more"`
		URL     string `json:"url"`
	} `json:"lines"`
	Livemode bool        `json:"livemode"`
	Memo     interface{} `json:"memo"`
	Metadata struct {
	} `json:"metadata"`
	Number          string      `json:"number"`
	OutOfBandAmount interface{} `json:"out_of_band_amount"`
	Pdf             string      `json:"pdf"`
	Reason          interface{} `json:"reason"`
	Refund          interface{} `json:"refund"`
	Status          string      `json:"status"`
	Subtotal        int         `json:"subtotal"`
	TaxAmounts      []struct {
		Amount    int    `json:"amount"`
		Inclusive bool   `json:"inclusive"`
		TaxRate   string `json:"tax_rate"`
	} `json:"tax_amounts"`
	Total    int         `json:"total"`
	Type     string      `json:"type"`
	VoidedAt interface{} `json:"voided_at"`
}