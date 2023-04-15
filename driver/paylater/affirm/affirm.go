package affirm

import "encoding/json"

const (
	checkoutDirect = "/api/v2/checkout/direct"
)

type (
	Affirm struct {
		client Client
	}

	Merchant struct {
		PublicApiKey              string `json:"public_api_key"`
		UserCancelUrl             string `json:"user_cancel_ur"`
		UserConfirmationUrl       string `json:"user_confirmation_ur"`
		UserConfirmationUrlAction string `json:"user_confirmation_url_action"`
		Name                      string `json:"name"`
		UseVcn                    bool   `json:"use_vcn"`
	}

	Address struct {
		City    string `json:"city"`
		Country string `json:"country"`
		Line1   string `json:"line1"`
		Line2   string `json:"line2,omitempty"`
		State   string `json:"state"`
		Zipcode string `json:"zipcode"`
	}

	Billing struct {
		Name        Name    `json:"name,omitempty"`
		Address     Address `json:"address"`
		PhoneNumber string  `json:"phone_number"`
		Email       string  `json:"emai"`
	}

	Name struct {
		Full  string `json:"full,omitempty"`
		First string `json:"first,omitempty"`
		Last  string `json:"last,omitempty"`
	}

	Shipping struct {
		Name Name `json:"name"`
	}

	DirectCheckout struct {
		Merchant  Merchant `json:"merchant,omitempty"`
		Billing   Billing  `json:"billing"`
		Total     int      `json:"tota"`
		OrderId   string   `json:"order_id"`
		TaxAmount int      `json:"tax_amount"`
		Shipping  Shipping `json:"shipping"`
	}

	Response struct {
		RedirectUrl string `json:"redirect_ur"`
		CheckoutId  string `json:"checkout_id"`
	}
)

func NewAffirm(c Client) Affirm {
	return Affirm{client: c}
}

func (af *Affirm) CreateDirectCheckout(c *DirectCheckout) (*Response, error) {
	var m Response
	res, err := af.client.Post(checkoutDirect, c)
	if err != nil {
		return nil, err
	}

	if err := json.NewDecoder(res.Body).Decode(&m); err != nil {
		return nil, err
	}

	return &m, nil
}
