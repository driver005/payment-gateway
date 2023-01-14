package models

import "github.com/driver005/gateway/core"

type PaymentMethod struct {
	core.Model

	BillingAddress *Address  `json:"billing_address" database:"foreignKey:id"`
	Card           *Card     `json:"card" database:"foreignKey:id"`
	Customer       *Customer `json:"customer" database:"foreignKey:id"`
	Type           string    `json:"type"`
}
