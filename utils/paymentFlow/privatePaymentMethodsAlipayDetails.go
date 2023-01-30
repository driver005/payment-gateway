package paymentFlow

import "github.com/driver005/gateway/core"

// PaymentFlowsPrivatePaymentMethodsAlipayDetails
type PaymentFlowsPrivatePaymentMethodsAlipayDetails struct {
	core.Model

	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	BuyerId string `json:"buyer_id,omitempty"`
	// Uniquely identifies this particular Alipay account. You can use this attribute to check whether two Alipay accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Transaction ID of this particular Alipay transaction.
	TransactionId string `json:"transaction_id,omitempty"`
}
