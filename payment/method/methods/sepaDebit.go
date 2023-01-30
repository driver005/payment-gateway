package methods

import "github.com/driver005/gateway/core"

type PaymentMethodSepaDebit struct {
	core.Model

	BankCode    string `json:"bank_code,omitempty"`
	BranchCode  string `json:"branch_code,omitempty"`
	Country     string `json:"country,omitempty"`
	Fingerprint string `json:"fingerprint,omitempty"`
	Last4       string `json:"last4,omitempty"`
}
