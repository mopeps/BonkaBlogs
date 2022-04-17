package data

import "time"

type Blog struct {
	ID	int64 `json:"id"`
	CreatedAt	time.Time `json:"-"`
	Title	string `json:"title"`
	Tags	[]string `json:"tags,omitempty"`
	Version	int32 `json:"version"`
}
