package methods

import "github.com/driver005/gateway/core"

type PaymentMethodLink struct {
	core.Model

	Email           string `json:"email,omitempty"`
	PersistentToken string `json:"persistent_token,omitempty"`
}
