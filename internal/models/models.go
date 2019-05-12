package models

import "time"

type CarOrder struct {
	ID      int `gorm:"primary_key"`
	Message string
	Offset  int
	Time    time.Time
}

type User struct {
	ID            int    `gorm:"primary_key"`
	OpenId        string
	NickName      string
	Sex           int
	Country       string
	Province      string
	City          string
	HeadImgUrl    string
	UnionId       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     time.Time
	LastLoginTime time.Time
	LastLoginIp   string
}
