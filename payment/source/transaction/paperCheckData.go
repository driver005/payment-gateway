package transaction

// SourceTransactionPaperCheckData
type SourceTransactionPaperCheckData struct {
	// Time at which the deposited funds will be available for use. Measured in seconds since the Unix epoch.
	AvailableAt *string `json:"available_at,omitempty"`
	// Comma-separated list of invoice IDs associated with the paper check.
	Invoices *string `json:"invoices,omitempty"`
}
