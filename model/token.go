package model

import "time"

type Token struct {
	Id          int
	AdminId     int
	AccessToken string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (t *Token) IsExpired() bool {
	return t.CreatedAt.Before(time.Now().Add(-time.Hour))
}
