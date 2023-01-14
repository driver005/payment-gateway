package models

// Application
type Application struct {
	// Unique identifier for the object.
	Id string `json:"id"`
	// The name of the application.
	Name string `json:"name,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
