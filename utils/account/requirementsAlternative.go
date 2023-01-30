package account

type AccountRequirementsAlternative struct {
	AlternativeFieldsDue []string `json:"alternative_fields_due"`
	OriginalFieldsDue    []string `json:"original_fields_due"`
}
