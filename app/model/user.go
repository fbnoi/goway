package model

import "flynoob/goway"

type UserType int

const (
	WECHAT_USER = UserType(1)
	PHONE_USER  = UserType(2)
)

type User struct {
	ID       int
	Client   *goway.Client
	NickName string
	Type     UserType

	RankScore int
}
