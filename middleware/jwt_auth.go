package middleware

import (
	"errors"
	"pocket-message/configs"
	"pocket-message/dto"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetToken(uuid uuid.UUID, username string) (string, error) {

	claims := jwt.MapClaims{}
	claims["uuid"] = uuid
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(configs.TokenSecret))
}

func DecodeJWT(ctx echo.Context) (dto.Token, error) {
	var t dto.Token

	auth := ctx.Request().Header.Get("Authorization")
	if auth == "" {
		return dto.Token{}, errors.New("authorization header not found")
	}

	splitToken := strings.Split(auth, "Bearer ")
	auth = splitToken[1]

	token, err := jwt.ParseWithClaims(auth, &dto.Token{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.TokenSecret), nil
	})
	if err != nil {
		return dto.Token{}, errors.New("token is wrong or expired")
	}

	if claims, ok := token.Claims.(*dto.Token); ok && token.Valid {
		t.UUID = claims.UUID
		t.Username = claims.Username
	}

	return t, nil
}

