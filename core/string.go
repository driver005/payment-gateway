package core

type String struct {
	N   string `json:"n,omitempty"`
	Lt  string `json:"lt,omitempty"`
	Gt  string `json:"gt,omitempty"`
	Gte string `json:"gte,omitempty"`
	Lte string `json:"lte,omitempty"`
}
