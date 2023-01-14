package next

// PaymentIntentNextActionKonbiniStores
type PaymentIntentNextActionKonbiniStores struct {
	Familymart PaymentIntentNextActionKonbiniStoresFamilymart `json:"familymart,omitempty"`
	Lawson     PaymentIntentNextActionKonbiniStoresLawson     `json:"lawson,omitempty"`
	Ministop   PaymentIntentNextActionKonbiniStoresMinistop   `json:"ministop,omitempty"`
	Seicomart  PaymentIntentNextActionKonbiniStoresSeicomart  `json:"seicomart,omitempty"`
}
