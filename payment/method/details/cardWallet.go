package details

// PaymentMethodDetailsCardWallet
type PaymentMethodDetailsCardWallet struct {
	//
	AmexExpressCheckout map[string]interface{} `json:"amex_express_checkout,omitempty"`
	//
	ApplePay map[string]interface{} `json:"apple_pay,omitempty"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string `json:"dynamic_last4,omitempty"`
	//
	GooglePay  map[string]interface{}                    `json:"google_pay,omitempty"`
	Masterpass *PaymentMethodDetailsCardWalletMasterpass `json:"masterpass,omitempty"`
	//
	SamsungPay map[string]interface{} `json:"samsung_pay,omitempty"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, or `visa_checkout`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         string                                      `json:"type"`
	VisaCheckout *PaymentMethodDetailsCardWalletVisaCheckout `json:"visa_checkout,omitempty"`
}
