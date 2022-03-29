package models

type Source struct {
	ID                string `json:"id"`
	Object            string `json:"object"`
	AchCreditTransfer struct {
		AccountNumber string `json:"account_number"`
		RoutingNumber string `json:"routing_number"`
		Fingerprint   string `json:"fingerprint"`
		BankName      string `json:"bank_name"`
		SwiftCode     string `json:"swift_code"`
	} `json:"ach_credit_transfer"`
	Amount       interface{} `json:"amount"`
	ClientSecret string      `json:"client_secret"`
	Created      int         `json:"created"`
	Currency     string      `json:"currency"`
	Flow         string      `json:"flow"`
	Livemode     bool        `json:"livemode"`
	Metadata     struct {
	} `json:"metadata"`
	Owner struct {
		Address         interface{} `json:"address"`
		Email           string      `json:"email"`
		Name            interface{} `json:"name"`
		Phone           interface{} `json:"phone"`
		VerifiedAddress interface{} `json:"verified_address"`
		VerifiedEmail   interface{} `json:"verified_email"`
		VerifiedName    interface{} `json:"verified_name"`
		VerifiedPhone   interface{} `json:"verified_phone"`
	} `json:"owner"`
	Receiver struct {
		Address                string `json:"address"`
		AmountCharged          int    `json:"amount_charged"`
		AmountReceived         int    `json:"amount_received"`
		AmountReturned         int    `json:"amount_returned"`
		RefundAttributesMethod string `json:"refund_attributes_method"`
		RefundAttributesStatus string `json:"refund_attributes_status"`
	} `json:"receiver"`
	StatementDescriptor interface{} `json:"statement_descriptor"`
	Status              string      `json:"status"`
	Type                string      `json:"type"`
	Usage               string      `json:"usage"`
}