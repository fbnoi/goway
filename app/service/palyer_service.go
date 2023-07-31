package service

import (
	"flynoob/goway"
	"flynoob/goway/app/model"
	pb "flynoob/goway/protobuf"
	"math"
	"sync"
	"time"
)

type queueNode struct {
	user *model.User
	prev *queueNode
	next *queueNode
}

type matchingQueue struct {
	head *queueNode
	end  *queueNode
}

func (mq *matchingQueue) append(user *model.User) {
	if user != nil {
		node := &queueNode{user: user, prev: mq.end}
		mq.end = node
		if mq.head == nil {
			mq.head = node
		}
	}
}

func (mq *matchingQueue) contain(user *model.User) bool {
	contain := false
	mq.walk(func(u *model.User) bool {
		if u == user {
			contain = true

			return false
		}
		return true
	})

	return contain
}

func (mq *matchingQueue) walk(fn func(user *model.User) bool) {
	cursor := mq.head
	for cursor != nil {
		if !fn(cursor.user) {
			break
		}
		cursor = cursor.next
	}
}

func (mq *matchingQueue) remove(user *model.User) bool {
	cursor := mq.head
	for cursor != nil {
		if user == cursor.user {
			prev := cursor.prev
			next := cursor.next
			if prev == nil && next == nil {
				mq.head = nil
				mq.end = nil
			} else if prev == nil {
				mq.head = next
				next.prev = prev
			} else if next == nil {
				mq.end = prev
				prev.next = nil
			} else {
				prev.next = next
				next.prev = prev
			}

			return true
		}
	}
	return false
}

type PlayerService struct {
	matchingUsers *matchingQueue
	mux           sync.Mutex
}

func (gs *PlayerService) OnStartMatching(client *goway.Client, frame *pb.Frame) {
	iUser, ok := client.Get("user")
	if ok {
		user, ok := iUser.(*model.User)
		if ok {
			gs.doMatching(user, client)
			return
		}
	}

	// frame := internal.GetFrame()
	// MatchingError := &pb.MatchingError{Code: 0, Message: ""}
	// frame
}

func (gs *PlayerService) doMatching(user *model.User, client *goway.Client) {
	gs.mux.Lock()
	defer gs.mux.Unlock()
	matched := false
	gs.matchingUsers.walk(func(waitingUser *model.User) bool {
		if match(waitingUser, user) {
			playground := gs.newPlayground(waitingUser, user)
			waitingUser.Client.Set("playground", playground)
			user.Client.Set("playground", playground)

			matched = true
		}
		return !matched
	})
	if !matched && !gs.matchingUsers.contain(user) {
		gs.matchingUsers.append(user)
	} else if matched {
		gs.matchingUsers.remove(user)
	}
}

func (gs *PlayerService) newPlayground(u1, u2 *model.User) *model.Playground {

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

func match(u1, u2 *model.User) bool {
	return math.Abs(float64(u1.RankScore-u2.RankScore)) < 10
}
