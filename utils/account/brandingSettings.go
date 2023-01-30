package account

import "github.com/driver005/gateway/internal/file"

type AccountBrandingSettings struct {
	Icon           file.File `json:"icon,omitempty"`
	Logo           file.File `json:"logo,omitempty"`
	PrimaryColor   string    `json:"primary_color,omitempty"`
	SecondaryColor string    `json:"secondary_color,omitempty"`
}
