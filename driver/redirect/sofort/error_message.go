package sofort

type errorMessage struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}
