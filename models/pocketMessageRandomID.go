package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PocketMessageRandomID struct {
	gorm.Model
	RandomID          string    `json:"random_id" form:"random_id"`
	Visit             int       `json:"visit"`
	PocketMessageUUID uuid.UUID `json:"pocket_message_uuid" form:"pocket_message_uuid" gorm:"type:VARCHAR(191)"`
}

type Tabler interface {
	TableName() string
}

func (PocketMessageRandomID) TableName() string {
	return "pocket_message_random_id"
}
