package account

type AccountPaymentsSettings struct {
	StatementDescriptor            string `json:"statement_descriptor,omitempty"`
	StatementDescriptorKana        string `json:"statement_descriptor_kana,omitempty"`
	StatementDescriptorKanji       string `json:"statement_descriptor_kanji,omitempty"`
	StatementDescriptorPrefixKana  string `json:"statement_descriptor_prefix_kana,omitempty"`
	StatementDescriptorPrefixKanji string `json:"statement_descriptor_prefix_kanji,omitempty"`
}
