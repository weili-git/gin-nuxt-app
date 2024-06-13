package models

import "time"

type Comment struct {
	ID      uint      `json:"id" gorm:"primary_key"`
	PostID  uint      `json:"post_id"`
	UserID  uint      `json:"user_id"`
	Body    string    `json:"body"`
	PubDate time.Time `json:"publish_date"`
	Likes   uint      `json:"likes"`
}
