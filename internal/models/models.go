package models

import "time"

type CarOrder struct {
	ID      int `gorm:"primary_key"`
	Message string
	Offset  int
	Time    time.Time
}

type BaseModel struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SmmsFile struct {
	BaseModel
	FileName  string `json:"filename"`
	StoreName string `json:"storename"`
	Size      int    `json:"size"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Hash      string `json:"hash"`
	Delete    string `json:"delete"`
	Url       string `json:"url"`
	Path      string `json:"path"`
}

type User struct {
	ID            int `gorm:"primary_key"`
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
