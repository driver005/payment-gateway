package models

type PaymentMethod struct {
	ID             string `json:"id"`
	Object         string `json:"object"`
	BillingDetails struct {
		Address struct {
			City       interface{} `json:"city"`
			Country    interface{} `json:"country"`
			Line1      interface{} `json:"line1"`
			Line2      interface{} `json:"line2"`
			PostalCode interface{} `json:"postal_code"`
			State      interface{} `json:"state"`
		} `json:"address"`
		Email string      `json:"email"`
		Name  interface{} `json:"name"`
		Phone string      `json:"phone"`
	} `json:"billing_details"`
	Card struct {
		Brand  string `json:"brand"`
		Checks struct {
			AddressLine1Check      interface{} `json:"address_line1_check"`
			AddressPostalCodeCheck interface{} `json:"address_postal_code_check"`
			CvcCheck               string      `json:"cvc_check"`
		} `json:"checks"`
		Country       string      `json:"country"`
		ExpMonth      int         `json:"exp_month"`
		ExpYear       int         `json:"exp_year"`
		Fingerprint   string      `json:"fingerprint"`
		Funding       string      `json:"funding"`
		GeneratedFrom interface{} `json:"generated_from"`
		Last4         string      `json:"last4"`
		Networks      struct {
			Available []string    `json:"available"`
			Preferred interface{} `json:"preferred"`
		} `json:"networks"`
		ThreeDSecureUsage struct {
			Supported bool `json:"supported"`
		} `json:"three_d_secure_usage"`
		Wallet interface{} `json:"wallet"`
	} `json:"card"`
	Created  int         `json:"created"`
	Customer interface{} `json:"customer"`
	Livemode bool        `json:"livemode"`
	Metadata struct {
		OrderID string `json:"order_id"`
	} `json:"metadata"`
	Type string `json:"type"`
}
