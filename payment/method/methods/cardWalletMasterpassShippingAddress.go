package methods

import "github.com/driver005/gateway/models"

// PaymentMethodCardWalletMasterpassShippingAddress Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
type PaymentMethodCardWalletMasterpassShippingAddress struct {
	Address *models.Address
}
