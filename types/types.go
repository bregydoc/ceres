package types

import (
	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `xml:"email"`

	Password string `json:"password"`
	Boss     string `json:"boss"`
}

type Admin struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `xml:"email"`

	Password string `json:"password"`
	Boss     string `json:"boss"`
}

type Team struct {
	ID    uuid.UUID `json:"id"`
	Name  string    `json:"name"`
	Email string    `xml:"email"`

	Password string `json:"password"`
	Boss     string `json:"boss"`
}
