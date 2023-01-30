package details

import "github.com/driver005/gateway/core"

// PaymentMethodDetailsWechatPay
type PaymentMethodDetailsWechatPay struct {
	core.Model

	// Uniquely identifies this particular WeChat Pay account. You can use this attribute to check whether two WeChat accounts are the same.
	Fingerprint string `json:"fingerprint,omitempty"`
	// Transaction ID of this particular WeChat Pay transaction.
	TransactionId string `json:"transaction_id,omitempty"`
}
