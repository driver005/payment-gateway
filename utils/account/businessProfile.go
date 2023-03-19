package account

import (
	"github.com/driver005/gateway/core"
	"github.com/driver005/gateway/utils/address"
)

type AccountBusinessProfile struct {
	core.Model

	Mcc                string           `json:"mcc,omitempty"`
	Name               string           `json:"name,omitempty"`
	ProductDescription string           `json:"product_description,omitempty"`
	SupportAddress     *address.Address `json:"support_address,omitempty" database:"foreignKey:id" swaggertype:"primitive,string" format:"uuid"`
	SupportEmail       string           `json:"support_email,omitempty"`
	SupportPhone       string           `json:"support_phone,omitempty"`
	SupportUrl         string           `json:"support_url,omitempty"`
	Url                string           `json:"url,omitempty"`
}
