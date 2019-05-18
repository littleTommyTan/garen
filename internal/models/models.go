package models

import "time"

type SmmsResponce struct {
	Code string `json:"code"`
	Data struct {
		FileName  string `json:"filename"`
		StoreName string `json:"storename"`
		Size      int    `json:"size"`
		Width     int    `json:"width"`
		Height    int    `json:"height"`
		Hash      string `json:"hash"`
		Delete    string `json:"delete"`
		Url       string `json:"url"`
		Path      string `json:"path"`
		Msg       string `json:"msg"`
	} `json:"data,omitempty"`
	Msg string `json:"msg,omitempty"`
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
