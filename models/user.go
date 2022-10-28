package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username      string `json:"username" form:"username"`
	Password      string `json:"password" form:"passwrd"`
	PocketMessage []PocketMessage
}
