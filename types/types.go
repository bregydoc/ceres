package types

import (
	"github.com/kataras/iris"
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

// Post ...
type Post struct {
	ID          uuid.UUID
	Title       string
	Description string
}

// VirtualTeam ...
type VirtualTeam struct {
	ID     uuid.UUID
	Name   string
	MakeBy iris.Handler
}
