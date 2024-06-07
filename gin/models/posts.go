package models

import "time"

type Post struct {
	ID       uint      `json:"id" gorm:"primary_key"`
	Title    string    `json:"title"`
	Body     string    `json:"body"`
	PubDate  time.Time `json:"publish_date"`
	ModDate  time.Time `json:"modify_date"`
	Category string    `json:"category"`
	Tag      string    `json:"tag"`
}
