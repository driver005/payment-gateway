package file

import (
	"github.com/driver005/gateway/core"
)

// FileLink To share the contents of a `File` object with non-Stripe users, you can create a `FileLink`. `FileLink`s contain a URL that can be used to retrieve the contents of the file without authentication.
type FileLink struct {
	core.Model

	// Whether this link is already expired.
	Expired bool `json:"expired,omitempty"`
	// Time at which the link expires.
	ExpiresAt int   `json:"expires_at,omitempty"`
	File      *File `json:"file,omitempty" database:"foreignKey:id"`
	// The publicly accessible URL to download the file.
	Url string `json:"url,omitempty"`
}

// File This is an object representing a file hosted on Stripe's servers. The file may have been uploaded by yourself using the [create file](https://stripe.com/docs/api#create_file) request (for example, when uploading dispute evidence) or it may have been created by Stripe (for example, the results of a [Sigma scheduled query](#scheduled_queries)).  Related guide: [File Upload Guide](https://stripe.com/docs/file-upload).
type File struct {
	core.Model

	// The time at which the file expires and is no longer available in epoch seconds.
	ExpiresAt int `json:"expires_at,omitempty"`
	// A filename for the file, suitable for saving to a filesystem.
	Filename string    `json:"filename,omitempty"`
	Links    *FileLink `json:"links,omitempty" database:"foreignKey:id"`
	// The [purpose](https://stripe.com/docs/file-upload#uploading-a-file) of the uploaded file.
	Purpose string `json:"purpose,omitempty"`
	// The size in bytes of the file object.
	Size int `json:"size,omitempty"`
	// A user friendly title for the document.
	Title string `json:"title,omitempty"`
	// The type of the file returned (e.g., `csv`, `pdf`, `jpg`, or `png`).
	Type string `json:"type,omitempty"`
	// The URL from which the file can be downloaded using your live secret API key.
	Url string `json:"url,omitempty"`
}
