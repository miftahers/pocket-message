package dto

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type Token struct {
	UUID     uuid.UUID `json:"uuid" form:"uuid"`
	Username string    `json:"username" form:"username"`
	jwt.StandardClaims
}
