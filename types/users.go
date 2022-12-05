package types

import "github.com/driver005/gateway/core"

type FilterableUserProps struct {
	core.Filter
	Email     string `json:"email,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
}
