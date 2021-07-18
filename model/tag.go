package model

import "time"

type Tags struct {
	ArticleId int
	Tags      string
	CreateAt  time.Time
}
