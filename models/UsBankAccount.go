package models

// LinkedAccountOptionsUsBankAccount
type LinkedAccountOptionsUsBankAccount struct {
	// The list of permissions to request. The `payment_method` permission must be included.
	Permissions []string `json:"permissions,omitempty"`
	// For webview integrations only. Upon completing OAuth login in the native browser, the user will be redirected to this URL to return to your app.
	ReturnUrl *string `json:"return_url,omitempty"`
}
