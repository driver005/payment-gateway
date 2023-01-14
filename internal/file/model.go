package file

// FileLink To share the contents of a `File` object with non-Stripe users, you can create a `FileLink`. `FileLink`s contain a URL that can be used to retrieve the contents of the file without authentication.
type FileLink struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// Whether this link is already expired.
	Expired bool `json:"expired"`
	// Time at which the link expires.
	ExpiresAt int   `json:"expires_at,omitempty"`
	File      *File `json:"file"`
	// Unique identifier for the object.
	Id string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The publicly accessible URL to download the file.
	Url string `json:"url,omitempty"`
}

// File This is an object representing a file hosted on Stripe's servers. The file may have been uploaded by yourself using the [create file](https://stripe.com/docs/api#create_file) request (for example, when uploading dispute evidence) or it may have been created by Stripe (for example, the results of a [Sigma scheduled query](#scheduled_queries)).  Related guide: [File Upload Guide](https://stripe.com/docs/file-upload).
type File struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int `json:"created"`
	// The time at which the file expires and is no longer available in epoch seconds.
	ExpiresAt int `json:"expires_at,omitempty"`
	// A filename for the file, suitable for saving to a filesystem.
	Filename string `json:"filename,omitempty"`
	// Unique identifier for the object.
	Id    string   `json:"id"`
	Links FileLink `json:"links,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The [purpose](https://stripe.com/docs/file-upload#uploading-a-file) of the uploaded file.
	Purpose string `json:"purpose"`
	// The size in bytes of the file object.
	Size int `json:"size"`
	// A user friendly title for the document.
	Title string `json:"title,omitempty"`
	// The type of the file returned (e.g., `csv`, `pdf`, `jpg`, or `png`).
	Type string `json:"type,omitempty"`
	// The URL from which the file can be downloaded using your live secret API key.
	Url string `json:"url,omitempty"`
}
