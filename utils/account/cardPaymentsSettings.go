package account

type AccountCardPaymentsSettings struct {
	DeclineOn                      *AccountDeclineChargeOn `json:"decline_on,omitempty"`
	StatementDescriptorPrefix      string                  `json:"statement_descriptor_prefix,omitempty"`
	StatementDescriptorPrefixKana  string                  `json:"statement_descriptor_prefix_kana,omitempty"`
	StatementDescriptorPrefixKanji string                  `json:"statement_descriptor_prefix_kanji,omitempty"`
}
