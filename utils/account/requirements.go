package account

type AccountRequirements struct {
	Alternatives        []AccountRequirementsAlternative `json:"alternatives,omitempty"`
	CurrentDeadline     int                              `json:"current_deadline,omitempty"`
	CurrentlyDue        []string                         `json:"currently_due,omitempty"`
	DisabledReason      string                           `json:"disabled_reason,omitempty"`
	Errors              []AccountRequirementsError       `json:"errors,omitempty"`
	EventuallyDue       []string                         `json:"eventually_due,omitempty"`
	PastDue             []string                         `json:"past_due,omitempty"`
	PendingVerification []string                         `json:"pending_verification,omitempty"`
}
