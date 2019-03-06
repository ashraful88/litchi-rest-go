package api

import "time"

// SampleEntity sample data to save
type SampleEntity struct {
	ID      string    `json:"id,omitempty"`
	Name    string    `json:"name,omitempty"`
	UserID  string    `json:"userid,omitempty" binding:"required"`
	Created time.Time `json:"created,omitempty"`
	Updated time.Time `json:"updated,omitempty"`
	Status  string    `json:"status,omitempty"`
}
