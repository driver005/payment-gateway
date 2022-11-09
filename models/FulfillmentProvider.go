package models

// Represents a fulfillment provider plugin and holds its installation status.
type FulfillmentProvider struct {
	// The id of the fulfillment provider as given by the plugin.
	Id *string `json:"id,omitempty"`

	// Whether the plugin is installed in the current version. Plugins that are no longer installed are not deleted by will have this field set to `false`.
	IsInstalled *bool `json:"is_installed,omitempty"`
}
