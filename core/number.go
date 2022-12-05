package core

type Number struct {
	N   int `json:"n,omitempty"`
	Lt  int `json:"lt,omitempty"`
	Gt  int `json:"gt,omitempty"`
	Gte int `json:"gte,omitempty"`
	Lte int `json:"lte,omitempty"`
}
