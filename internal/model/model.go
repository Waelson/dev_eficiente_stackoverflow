package model

import "time"

const PostStatusClosed = 2
const PostStatusOpened = 1

type Post struct {
	Id          int64    `json:"id"`
	Title       string   `json:"title" binding:"required,not_blank,max=150"`
	Description string   `json:"description" binding:"required,not_blank,max=1500"`
	Tags        []string `json:"tags" binding:"required,tags,duplicated"`
	User        string
	CreateAt    time.Time
	Status      int
}

func (a *Post) IsOpened() bool {
	return a.Status == PostStatusClosed
}

func (a *Post) IsClosed() bool {
	return a.Status == PostStatusClosed
}

type Answer struct {
	Id       int64  `json:"id"`
	PostId   int64  `json:"post_id"`
	Response string `json:"response" binding:"required,not_blank,min=20"`
	User     string
	CreateAt time.Time
}

type User struct {
	Id       int64  `json:"id"`
	Name     string `json:"name" binding:"required,not_blank,max=150"`
	Email    string `json:"email" binding:"required,not_blank,max=150"`
	Login    string `json:"login" binding:"required,not_blank,max=50"`
	Password string `json:"password" binding:"required,not_blank,min=8"`
	CreateAt time.Time
}

type Login struct {
	Login    string `json:"login" binding:"required,not_blank,max=50"`
	Password string `json:"password" binding:"required,not_blank,min=8"`
}
