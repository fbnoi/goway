package model

import "time"

type PlayerType int

const (
	PLAYER_WHITE = PlayerType(1)
	PLAYER_BLACK = PlayerType(2)
)

type Player struct {
	ID        int
	User      *User
	Type      PlayerType
	CreatedAt time.Time
}

func NewPlayer(user *User, typ PlayerType) *Player {
	return &Player{
		User:      user,
		Type:      typ,
		CreatedAt: time.Now(),
	}
}
