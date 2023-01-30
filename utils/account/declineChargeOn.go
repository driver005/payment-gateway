package account

type AccountDeclineChargeOn struct {
	AvsFailure bool `json:"avs_failure"`
	CvcFailure bool `json:"cvc_failure"`
}
