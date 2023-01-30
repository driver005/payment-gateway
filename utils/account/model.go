package account

import (
	"github.com/driver005/gateway/core"
)

// type Polymorphic struct {
// 	BankAccount bank.BankAccount
// 	Card        card.Card
// }

type Account struct {
	core.Model

	BusinessProfile *AccountBusinessProfile `json:"business_profile,omitempty" database:"foreignKey:id"`
	// BusinessType       string                               `json:"business_type,omitempty"`
	// Capabilities       *AccountCapabilities                 `json:"capabilities,omitempty"`
	// ChargesEnabled     bool                                `json:"charges_enabled,omitempty"`
	// Company            *LegalEntityCompany                  `json:"company,omitempty"`
	// Controller         *AccountUnificationAccountController `json:"controller,omitempty"`
	// Country            string                               `json:"country,omitempty"`
	// DefaultCurrency    string                               `json:"default_currency,omitempty"`
	// DetailsSubmitted   bool                                `json:"details_submitted,omitempty"`
	// Email              string                               `json:"email,omitempty"`
	// ExternalAccounts   []Polymorphic                        `json:"external_accounts,omitempty"`
	// FutureRequirements *AccountFutureRequirements           `json:"future_requirements,omitempty"`
	// Individual         *Person                              `json:"individual,omitempty"`
	// PayoutsEnabled     bool                                `json:"payouts_enabled,omitempty"`
	// Requirements       *AccountRequirements                 `json:"requirements,omitempty"`
	// Settings           AccountSettings                      `json:"settings,omitempty"`
	// TosAcceptance      *AccountTosAcceptance                `json:"tos_acceptance,omitempty"`
	// Type               string                               `json:"type,omitempty"`
}
