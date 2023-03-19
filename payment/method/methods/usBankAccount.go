package methods

import (
	"github.com/driver005/gateway/core"
	"github.com/lib/pq"
)

type PaymentMethodUsBankAccount struct {
	core.Model

	AccountHolderType           string         `json:"account_holder_type,omitempty"`
	AccountType                 string         `json:"account_type,omitempty"`
	BankName                    string         `json:"bank_name,omitempty"`
	FinancialConnectionsAccount string         `json:"financial_connections_account,omitempty"`
	Fingerprint                 string         `json:"fingerprint,omitempty"`
	Last4                       string         `json:"last4,omitempty"`
	PreferredNetwork            string         `json:"preferred_network,omitempty"`
	SupportedNetworks           pq.StringArray `json:"available_networks,omitempty" database:"type:varchar(64)[]"`
	RoutingNumber               string         `json:"routing_number,omitempty"`
}
