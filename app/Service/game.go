package service

import (
	"flynoob/goway"
	"flynoob/goway/app/model"
	pb "flynoob/goway/protobuf"
	"math"
	"time"
)

var waitingSet []*model.User

func OnStartGame(client *goway.Client, frame *pb.Frame) {
	if user := GetUserFromClient(client); user != nil {
		// for _, wu := range waitingSet {
		// }
		waitingSet = append(waitingSet, user)
	}
}

func match(u1, u2 *model.User) bool {
	return math.Abs(float64(u1.RankScore-u2.RankScore)) < 10
}

func newPlayground(u1, u2 *model.User) *model.Playground {

	var whitePlayer, blackPlayer *model.Player
	if u1.RankScore > u2.RankScore {
		whitePlayer = model.NewPlayer(u1, model.PLAYER_WHITE)
		blackPlayer = model.NewPlayer(u2, model.PLAYER_BLACK)
	} else {
		whitePlayer = model.NewPlayer(u2, model.PLAYER_WHITE)
		blackPlayer = model.NewPlayer(u1, model.PLAYER_BLACK)
	}

	return &model.Playground{
		WP:        whitePlayer,
		BP:        blackPlayer,
		CreatedAt: time.Now(),
		Status:    model.PLAYING,
	}
}
