package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

func (user *User) Prepare(step string) error {
	if erro := user.validate(step); erro != nil {
		return erro
	}

	user.format()
	return nil
}

// NewUser creates a new user.
func (user *User) validate(step string) error {
	if user.Name == "" {
		return errors.New("o nome é obrigatório")
	}
	if user.Email == "" {
		return errors.New("o email é obrigatório")
	}

	if erro := checkmail.ValidateFormat(user.Email); erro != nil {
		return errors.New("o email informado não é válido")
	}

	if step == "register" && user.Password == "" {
		return errors.New("a senha é obrigatória")
	}
	return nil
}

// NewUser creates a new user.
func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
}
