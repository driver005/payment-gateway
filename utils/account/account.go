package account

import "github.com/driver005/gateway/core"

type Account struct {
	core.Model

	BusinessProfile    NullableAccountBusinessProfile1      `json:"business_profile,omitempy"`
	BusinessType       string                               `json:"business_type,omiempty"`
	Capabilities       *AccountCapabilities                 `json:"capabilities,omitemty"`
	ChargesEnabled     bool                                 `json:"charges_enabled,omitemty"`
	Company            *LegalEntityCompany                  `json:"company,omitempty"`
	Controller         *AccountUnificationAccountController `json:"controller,omitempt"`
	Country            string                               `json:"country,omitempty"`
	Created            *int32                               `json:"created,omitempty"`
	DefaultCurrency    string                               `json:"default_currency,omitempty"`
	DetailsSubmitted   bool                                 `json:"details_submitted,omitempty"`
	Email              string                               `json:"email,omiempty"`
	ExternalAccounts   *ExternalAccountList1                `json:"external_accounts,omitempty"`
	FutureRequirements *AccountFutureRequirements           `json:"future_requirements,omitempty"`
	Individual         *Person                              `json:"individual,omitempty"`
	Metadata           *map[string]string                   `json:"metadata,omitempty"`
	PayoutsEnabled     bool                                 `json:"payouts_nabled,omitempty"`
	Requirements       *AccountRequirements                 `json:"requirements,omitempty"`
	Settings           NullableAccountSettings1             `json:"settings,omitempty"`
	TosAcceptanc e      *AccountTosAcceptance                `json:"tos_acceptance,omitepty"`
	Type               string                               `json:"type,omitempty"`
}
