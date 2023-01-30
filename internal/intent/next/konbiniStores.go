package next

import "github.com/driver005/gateway/core"

// PaymentIntentNextActionKonbiniStores
type PaymentIntentNextActionKonbiniStores struct {
	core.Model

	Familymart PaymentIntentNextActionKonbiniStoresFamilymart `json:"familymart,omitempty" database:"foreignKey:id"`
	Lawson     PaymentIntentNextActionKonbiniStoresLawson     `json:"lawson,omitempty" database:"foreignKey:id"`
	Ministop   PaymentIntentNextActionKonbiniStoresMinistop   `json:"ministop,omitempty" database:"foreignKey:id"`
	Seicomart  PaymentIntentNextActionKonbiniStoresSeicomart  `json:"seicomart,omitempty" database:"foreignKey:id"`
}
