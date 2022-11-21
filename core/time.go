package core

import "time"

type Time struct {
	Lt  time.Time `json:"lt"`
	Gt  time.Time `json:"gt"`
	Gte time.Time `json:"gte"`
	Lte time.Time `json:"lte"`
}
