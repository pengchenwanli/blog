package model

import "time"

type Article struct {
	Id       int
	Author   string
	Abstract string
	Content  string
	Title    string
	Creator  string  //int
	Updater  string  //int
	CreateAt time.Time
	UpdateAt time.Time
	DeleteAt *time.Time
}
