package account

import "github.com/lib/pq"

type AccountRequirementsAlternative struct {
	AlternativeFieldsDue pq.StringArray `json:"alternative_fields_due"`
	OriginalFieldsDue    pq.StringArray `json:"original_fields_due"`
}
