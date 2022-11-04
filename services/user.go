package services

import (
	"errors"
	"fmt"
	"net/http"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type UserServices interface {
	SignUp(echo.Context) error
	Login(echo.Context) (dto.Login, error)
	UpdateUsername(echo.Context) error
	UpdatePassword(echo.Context) error
}

type userServices struct {
	repositories.Database
}

// TODO SignUp Unit Test
func (s *userServices) SignUp(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return err
	}

	if u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error - password should not be empty",
		})
	}

	uuid := uuid.New()
	u.UUID = uuid

	err = s.Database.SaveNewUser(u)
	if err != nil {
		return err
	}

	return nil
}

// TODO Login Unit Test
func (s *userServices) Login(c echo.Context) (dto.Login, error) {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return dto.Login{}, err
	}
	if u.Username == "" {
		return dto.Login{}, errors.New("error, username should not be empty")
	}
	if u.Password == "" {
		return dto.Login{}, errors.New("error, password should not be empty")
	}

	user, err := s.Database.Login(u)
	if err != nil {
		return dto.Login{}, err
	}

	token, err := middleware.GetToken(user.UUID, user.Username)
	if err != nil {
		return dto.Login{}, err
	}

	var result dto.Login
	result.Username = user.Username
	result.Token = token

	return result, nil
}

// TODO UpdateUsername Unit Test
func (s *userServices) UpdateUsername(c echo.Context) error {
	var u models.User
	c.Bind(&u)
	if u.Username == "" {
		return errors.New("error, username should not be empty")
	}
	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return err
	}
	u.UUID = t.UUID

	fmt.Println(u)

	err = s.Database.UpdateUsername(u)
	if err != nil {
		return err
	}

	return nil
}

// TODO UpdatePassword Unit Test
func (s *userServices) UpdatePassword(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return err
	}

	if u.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error, username should not be empty",
		})
	}
	if u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": "error, password should not be empty",
		})
	}

	err = s.Database.UpdatePassword(u)
	if err != nil {
		return err
	}

	return nil
}