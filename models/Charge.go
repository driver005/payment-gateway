package models

type Charge struct {
	ID                   string      `json:"id"`
	Object               string      `json:"object"`
	Amount               int         `json:"amount"`
	AmountCaptured       int         `json:"amount_captured"`
	AmountRefunded       int         `json:"amount_refunded"`
	Application          interface{} `json:"application"`
	ApplicationFee       interface{} `json:"application_fee"`
	ApplicationFeeAmount interface{} `json:"application_fee_amount"`
	BalanceTransaction   string      `json:"balance_transaction"`
	BillingDetails       struct {
		Address struct {
			City       interface{} `json:"city"`
			Country    interface{} `json:"country"`
			Line1      interface{} `json:"line1"`
			Line2      interface{} `json:"line2"`
			PostalCode interface{} `json:"postal_code"`
			State      interface{} `json:"state"`
		} `json:"address"`
		Email interface{} `json:"email"`
		Name  string      `json:"name"`
		Phone interface{} `json:"phone"`
	} `json:"billing_details"`
	CalculatedStatementDescriptor interface{} `json:"calculated_statement_descriptor"`
	Captured                      bool        `json:"captured"`
	Created                       int         `json:"created"`
	Currency                      string      `json:"currency"`
	Customer                      interface{} `json:"customer"`
	Description                   string      `json:"description"`
	Disputed                      bool        `json:"disputed"`
	FailureBalanceTransaction     interface{} `json:"failure_balance_transaction"`
	FailureCode                   interface{} `json:"failure_code"`
	FailureMessage                interface{} `json:"failure_message"`
	FraudDetails                  struct {
	} `json:"fraud_details"`
	Invoice  interface{} `json:"invoice"`
	Livemode bool        `json:"livemode"`
	Metadata struct {
	} `json:"metadata"`
	OnBehalfOf           interface{} `json:"on_behalf_of"`
	Order                interface{} `json:"order"`
	Outcome              interface{} `json:"outcome"`
	Paid                 bool        `json:"paid"`
	PaymentIntent        interface{} `json:"payment_intent"`
	PaymentMethod        string      `json:"payment_method"`
	PaymentMethodDetails struct {
		Card struct {
			Brand  string `json:"brand"`
			Checks struct {
				AddressLine1Check      interface{} `json:"address_line1_check"`
				AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
				CvcCheck               string      `json:"cvc_check"`
			} `json:"checks"`
			Country      string      `json:"country"`
			ExpMonth     int         `json:"exp_month"`
			ExpYear      int         `json:"exp_year"`
			Fingerprint  string      `json:"fingerprint"`
			Funding      string      `json:"funding"`
			Installments interface{} `json:"installments"`
			Last4        string      `json:"last4"`
			Mandate      interface{} `json:"mandate"`
			Moto         interface{} `json:"moto"`
			Network      string      `json:"network"`
			ThreeDSecure interface{} `json:"three_d_secure"`
			Wallet       interface{} `json:"wallet"`
		} `json:"card"`
		Type string `json:"type"`
	} `json:"payment_method_details"`
	ReceiptEmail  interface{} `json:"receipt_email"`
	ReceiptNumber interface{} `json:"receipt_number"`
	ReceiptURL    string      `json:"receipt_url"`
	Redaction     interface{} `json:"redaction"`
	Refunded      bool        `json:"refunded"`
	Refunds       struct {
		Object  string        `json:"object"`
		Data    []interface{} `json:"data"`
		HasMore bool          `json:"has_more"`
		URL     string        `json:"url"`
	} `json:"refunds"`
	Review                    interface{} `json:"review"`
	Shipping                  interface{} `json:"shipping"`
	SourceTransfer            interface{} `json:"source_transfer"`
	StatementDescriptor       interface{} `json:"statement_descriptor"`
	StatementDescriptorSuffix interface{} `json:"statement_descriptor_suffix"`
	Status                    string      `json:"status"`
	TransferData              interface{} `json:"transfer_data"`
	TransferGroup             interface{} `json:"transfer_group"`
}