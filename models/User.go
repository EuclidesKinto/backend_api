package models

import "time"

// User represents a user of the application.
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}
