package types

import (
	"time"

	uuid "github.com/satori/go.uuid"
)

// User ...
type User struct {
	ID           uuid.UUID      `json:"id" form:"id"`
	Name         string         `json:"name" form:"name"`
	Email        string         `json:"email" form:"email"`
	Birthday     time.Time      `json:"birthday" form:"birthday"`
	Password     string         `json:"password" form:"password"`
	Username     string         `json:"username" form:"username"`
	OldTeams     []*VirtualTeam `json:"old_teams" form:"old_teams"`
	PlayingTeams []*VirtualTeam `json:"playing_teams" form:"playing_teams"`
	SavedTeams   []*VirtualTeam `json:"saved_teams" form:"saved_teams"`
}
