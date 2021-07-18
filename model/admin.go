package model

import "time"

//此处的admin代表的是用户
type Admin struct {
	Id          int
	AccountName string
	Password    string
	CreateAt    time.Time
	UpdateAt    time.Time
	DeleteAT    *time.Time
}
