package model

import "time"

type LoginReq struct {
	UserName string
	Password string
}

type LoginResp struct {
	Token        string
	RefreshToken string
	ExpireAt     time.Time
}
