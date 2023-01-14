package methods

import "github.com/driver005/gateway/models"

// PaymentMethodCardWalletMasterpassBillingAddress Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
type PaymentMethodCardWalletMasterpassBillingAddress struct {
	Address *models.Address
}
