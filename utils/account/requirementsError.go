package account

type AccountRequirementsError struct {
	Code        string `json:"code"`
	Reason      string `json:"reason"`
	Requirement string `json:"requirement"`
}
