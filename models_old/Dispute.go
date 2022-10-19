package models

type Disputes struct {
	ID                  string        `json:"id"`
	Object              string        `json:"object"`
	Amount              int           `json:"amount"`
	BalanceTransactions []interface{} `json:"balance_transactions"`
	Charge              string        `json:"charge"`
	Created             int           `json:"created"`
	Currency            string        `json:"currency"`
	Evidence            struct {
		AccessActivityLog            interface{} `json:"access_activity_log"`
		BillingAddress               interface{} `json:"billing_address"`
		CancellationPolicy           interface{} `json:"cancellation_policy"`
		CancellationPolicyDisclosure interface{} `json:"cancellation_policy_disclosure"`
		CancellationRebuttal         interface{} `json:"cancellation_rebuttal"`
		CustomerCommunication        interface{} `json:"customer_communication"`
		CustomerEmailAddress         interface{} `json:"customer_email_address"`
		CustomerName                 interface{} `json:"customer_name"`
		CustomerPurchaseIP           interface{} `json:"customer_purchase_ip"`
		CustomerSignature            interface{} `json:"customer_signature"`
		DuplicateChargeDocumentation interface{} `json:"duplicate_charge_documentation"`
		DuplicateChargeExplanation   interface{} `json:"duplicate_charge_explanation"`
		DuplicateChargeID            interface{} `json:"duplicate_charge_id"`
		ProductDescription           interface{} `json:"product_description"`
		Receipt                      interface{} `json:"receipt"`
		RefundPolicy                 interface{} `json:"refund_policy"`
		RefundPolicyDisclosure       interface{} `json:"refund_policy_disclosure"`
		RefundRefusalExplanation     interface{} `json:"refund_refusal_explanation"`
		ServiceDate                  interface{} `json:"service_date"`
		ServiceDocumentation         interface{} `json:"service_documentation"`
		ShippingAddress              interface{} `json:"shipping_address"`
		ShippingCarrier              interface{} `json:"shipping_carrier"`
		ShippingDate                 interface{} `json:"shipping_date"`
		ShippingDocumentation        interface{} `json:"shipping_documentation"`
		ShippingTrackingNumber       interface{} `json:"shipping_tracking_number"`
		UncategorizedFile            interface{} `json:"uncategorized_file"`
		UncategorizedText            interface{} `json:"uncategorized_text"`
	} `json:"evidence"`
	EvidenceDetails struct {
		DueBy           int  `json:"due_by"`
		HasEvidence     bool `json:"has_evidence"`
		PastDue         bool `json:"past_due"`
		SubmissionCount int  `json:"submission_count"`
	} `json:"evidence_details"`
	IsChargeRefundable bool `json:"is_charge_refundable"`
	Livemode           bool `json:"livemode"`
	Metadata           struct {
	} `json:"metadata"`
	PaymentIntent interface{} `json:"payment_intent"`
	Reason        string      `json:"reason"`
	Status        string      `json:"status"`
}
