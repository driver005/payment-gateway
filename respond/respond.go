package respond

import "net/http"

type Respond interface {
	NewResponse(w http.ResponseWriter) *Response
}
