package account

import "github.com/lib/pq"

type AccountFutureRequirements struct {
	Alternatives        []AccountRequirementsAlternative `json:"alternatives,omitempty"`
	CurrentDeadline     int                              `json:"current_deadline,omitempty"`
	CurrentlyDue        pq.StringArray                   `json:"currently_due,omitempty"`
	DisabledReason      string                           `json:"disabled_reason,omitempty"`
	Errors              []AccountRequirementsError       `json:"errors,omitempty"`
	EventuallyDue       pq.StringArray                   `json:"eventually_due,omitempty"`
	PastDue             pq.StringArray                   `json:"past_due,omitempty"`
	PendingVerification pq.StringArray                   `json:"pending_verification,omitempty"`
}
