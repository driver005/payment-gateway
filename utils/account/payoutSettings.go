package account

// TransferSchedule
type TransferSchedule struct {
	// The number of days charges for the account will be held before being paid out.
	DelayDays int `json:"delay_days"`
	// How frequently funds will be paid out. One of `manual` (payouts only created via API call), `daily`, `weekly`, or `monthly`.
	Interval string `json:"interval"`
	// The day of the month funds will be paid out. Only shown if `interval` is monthly. Payouts scheduled between the 29th and 31st of the month are sent on the last day of shorter months.
	MonthlyAnchor int `json:"monthly_anchor,omitempty"`
	// The day of the week funds will be paid out, of the style 'monday', 'tuesday', etc. Only shown if `interval` is weekly.
	WeeklyAnchor string `json:"weekly_anchor,omitempty"`
}

type AccountPayoutSettings struct {
	DebitNegativeBalances bool             `json:"debit_negative_balances"`
	Schedule              TransferSchedule `json:"schedule"`
	StatementDescriptor   string           `json:"statement_descriptor,omitempty"`
}
