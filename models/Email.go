package models

// EmailSent
type EmailSent struct {
	// The timestamp when the email was sent.
	EmailSentAt int `json:"email_sent_at"`
	// The recipient's email address.
	EmailSentTo string `json:"email_sent_to"`
}
