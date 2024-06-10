package entity

import (
	"net"
	"time"
)

type AuthData struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type User struct {
	Id           uint32
	Login        string
	PasswordHash string
	CreatedAt    time.Time
	Secret       string
}

type Session struct {
	Id        string
	UId       uint32
	CreatedAt time.Time
	Ip        net.IP
}
