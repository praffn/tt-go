package trat

import "time"

type Entry struct {
	Start time.Time `json:"start"`
	End   time.Time `json:"end"`
	Note  string    `json:"note"`
}

type Entries struct {
	Entries []Entry `json:"entries"`
}
