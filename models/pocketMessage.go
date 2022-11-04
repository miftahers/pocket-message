package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PocketMessage struct {
	gorm.Model
	UUID     uuid.UUID `json:"uuid" gorm:"primaryKey"`
	Title    string    `json:"title" form:"title"`
	Content  string    `json:"content" form:"content"`
	UserUUID uuid.UUID `json:"user_uuid" form:"user_uuid"`
}

func (PocketMessage) TableName() string {
	return "pocket_messages"
}
