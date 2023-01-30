package account

type AccountCapabilityFutureRequirements struct {
	Alternatives        []AccountRequirementsAlternative `json:"alternatives,omitempty"`
	CurrentDeadline     int                              `json:"current_deadline,omitempty"`
	CurrentlyDue        []string                         `json:"currently_due"`
	DisabledReason      string                           `json:"disabled_reason,omitempty"`
	Errors              []AccountRequirementsError       `json:"errors"`
	EventuallyDue       []string                         `json:"eventually_due"`
	PastDue             []string                         `json:"past_due"`
	PendingVerification []string                         `json:"pending_verification"`
}
