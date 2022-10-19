package models

type Event struct {
	ID         string `json:"id"`
	Object     string `json:"object"`
	APIVersion string `json:"api_version"`
	Created    int    `json:"created"`
	Data       struct {
		Object struct {
			ID           string `json:"id"`
			Object       string `json:"object"`
			Amount       int    `json:"amount"`
			ClientSecret string `json:"client_secret"`
			Created      int    `json:"created"`
			Currency     string `json:"currency"`
			Flow         string `json:"flow"`
			Livemode     bool   `json:"livemode"`
			Metadata     struct {
			} `json:"metadata"`
			Owner struct {
				Address         interface{} `json:"address"`
				Email           interface{} `json:"email"`
				Name            interface{} `json:"name"`
				Phone           interface{} `json:"phone"`
				VerifiedAddress interface{} `json:"verified_address"`
				VerifiedEmail   interface{} `json:"verified_email"`
				VerifiedName    string      `json:"verified_name"`
				VerifiedPhone   interface{} `json:"verified_phone"`
			} `json:"owner"`
			Redirect struct {
				FailureReason interface{} `json:"failure_reason"`
				ReturnURL     string      `json:"return_url"`
				Status        string      `json:"status"`
				URL           string      `json:"url"`
			} `json:"redirect"`
			Sofort struct {
				Country             string      `json:"country"`
				BankCode            string      `json:"bank_code"`
				BankName            string      `json:"bank_name"`
				Bic                 string      `json:"bic"`
				IbanLast4           string      `json:"iban_last4"`
				StatementDescriptor interface{} `json:"statement_descriptor"`
				PreferredLanguage   interface{} `json:"preferred_language"`
			} `json:"sofort"`
			StatementDescriptor interface{} `json:"statement_descriptor"`
			Status              string      `json:"status"`
			Type                string      `json:"type"`
			Usage               string      `json:"usage"`
		} `json:"object"`
	} `json:"data"`
	Livemode        bool `json:"livemode"`
	PendingWebhooks int  `json:"pending_webhooks"`
	Request         struct {
		ID             interface{} `json:"id"`
		IdempotencyKey interface{} `json:"idempotency_key"`
	} `json:"request"`
	Type string `json:"type"`
}
