package models

import (
	"errors"
	"strings"
	"time"
)

// User represents a user of the application.
type User struct {
	ID        int64     `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created,omitempty"`
	UpdatedAt time.Time `json:"updated,omitempty"`
}

func (user *User) Prepare() error {
	if erro := user.validate(); erro != nil {
		return erro
	}

	user.format()
	return nil
}

// NewUser creates a new user.
func (user *User) validate() error {
	if user.Name == "" {
		return errors.New("O nome é obrigatório.")
	}
	if user.Email == "" {
		return errors.New("O email é obrigatório.")
	}
	if user.Password == "" {
		return errors.New("A senha é obrigatória.")
	}
	return nil
}

// NewUser creates a new user.
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}
