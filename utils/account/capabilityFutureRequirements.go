package account

import "github.com/lib/pq"

type AccountCapabilityFutureRequirements struct {
	Alternatives        []AccountRequirementsAlternative `json:"alternatives,omitempty"`
	CurrentDeadline     int                              `json:"current_deadline,omitempty"`
	CurrentlyDue        pq.StringArray                   `json:"currently_due"`
	DisabledReason      string                           `json:"disabled_reason,omitempty"`
	Errors              []AccountRequirementsError       `json:"errors"`
	EventuallyDue       pq.StringArray                   `json:"eventually_due"`
	PastDue             pq.StringArray                   `json:"past_due"`
	PendingVerification pq.StringArray                   `json:"pending_verification"`
}
