package model

import "time"

type Token struct {
	AuthID    int64
	Token     string
	ExpiredAt time.Time
}
