package core

import "time"

type Time struct {
	N   time.Time `json:"n,omitempty"`
	Lt  time.Time `json:"lt,omitempty"`
	Gt  time.Time `json:"gt,omitempty"`
	Gte time.Time `json:"gte,omitempty"`
	Lte time.Time `json:"lte,omitempty"`
}
