package models

import (
	"time"

	"github.com/driver005/gateway/core"
)

// Invite - Represents an invite
type Invite struct {
	core.Model

	// The email of the user being invited.
	UserEmail string `json:"user_email"`

	// The user's role.
	Role string `json:"role" database:"default:null"`

	// Whether the invite was accepted or not.
	Accepted bool `json:"accepted" database:"default:null"`

	// The token used to accept the invite.
	Token string `json:"token" database:"default:null"`

	// The date the invite expires at.
	ExporesAt time.Time `json:"expores_at" database:"default:null"`

	// An optional key-value map with additional details
	Metadata JSONB `json:"metadata" database:"default:null"`
}
