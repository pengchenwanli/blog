package model

import "time"

type Comment struct {
	Id        int
	UserId    int //int
	ArticleId int
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
}
