package models

import "gorm.io/gorm"

type URL struct {
	gorm.Model
	Url               string `json:"url" form:"url"`
	Pocket_message_id uint   `json:"pocket_message_id" form:"pocket_message_id"`
}
