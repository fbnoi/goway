package model

import "time"

type ResultType int

func (p ResultType) String() string {
	switch p {
	case WHITE_WIN:
		return "White player win"
	case BLACK_WIN:
		return "Black player win"
	case ABORT:
		return "Abort"
	case DRAWN:
		return "Drown"
	}
	return "-"
}

const (
	WHITE_WIN = ResultType(1)
	BLACK_WIN = ResultType(2)
	ABORT     = ResultType(3)
	DRAWN     = ResultType(4)
)

type PlaygroundStatus int

func (p PlaygroundStatus) String() string {
	switch p {
	case PLAYING:
		return "Playing"
	case FINISHED:
		return "Finished"
	}
	return "-"
}

const (
	PLAYING  = PlaygroundStatus(1)
	SUSPEND  = PlaygroundStatus(2)
	FINISHED = PlaygroundStatus(3)
)

type Playground struct {
	ID        int
	WP        *Player
	BP        *Player
	CreatedAt time.Time
	EndedAt   time.Time
	Status    PlaygroundStatus
	Result    ResultType

	Watchers []*User
}
