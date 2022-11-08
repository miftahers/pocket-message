package controllers

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/models"

	"github.com/labstack/echo/v4"
)

type MockUserServices struct{}

func (s *MockUserServices) SignUp(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return err
	}

	if u.Password == "" {
		return errors.New("error, password should not be empty")
	}

	return nil
}
func (s *MockUserServices) Login(c echo.Context) (dto.Login, error) {
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

	return dto.Login{
		Username: "Super",
		Token:    "Idol",
	}, nil
}
func (s *MockUserServices) UpdateUsername(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	if u.Username == "" {
		return errors.New("error, username should not be empty")
	}

	return nil
}
func (s *MockUserServices) UpdatePassword(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return err
	}

	if u.Username == "" {
		return errors.New("error, username should not be empty")
	}
	if u.Password == "" {
		return errors.New("error, password should not be empty")
	}

	return nil
}
