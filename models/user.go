package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UUID          uuid.UUID       `json:"uuid" gorm:"primaryKey;type=uuid"`
	Username      string          `json:"username" form:"username"`
	Password      string          `json:"password" form:"password"`
	PocketMessage []PocketMessage `gorm:"foreignKey:UserUUID;references:UUID;type:VARCHAR(191);constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (User) TableName() string {
	return "users"
}
