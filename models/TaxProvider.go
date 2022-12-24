package models

import "github.com/google/uuid"

// TaxProvider - The tax service used to calculate taxes
type TaxProvider struct {

	// The id of the tax provider as given by the plugin.
	Id uuid.UUID `json:"id database:primarykey"`

	// Whether the plugin is installed in the current version. Plugins that are no longer installed are not deleted by will have this field set to `false`.
	IsInstalled bool `json:"is_installed" database:"default:null"`
}
