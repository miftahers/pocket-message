package models

import "gorm.io/gorm"

type PocketMessage struct {
	gorm.Model
	Title   string `json:"title" form:"title"`
	Content string `json:"content" form:"content"`
	UserID  string `json:"user_id" form:"user_id"`
	URL     URL
}
