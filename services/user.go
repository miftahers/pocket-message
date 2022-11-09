package services

import (
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
)

func NewUserServices(db repositories.Database) UserServices {
	return &userServices{Database: db}
}

type UserServices interface {
	SignUp(user models.User) error
	Login(user models.User) (dto.Login, error)
	UpdateUsername(user models.User, token dto.Token) error
	UpdatePassword(user models.User) error
}

type userServices struct {
	repositories.Database
}

func (s *userServices) SignUp(u models.User) error {

	u.UUID = uuid.New()

	err := s.Database.SaveNewUser(u)
	if err != nil {
		return err
	}

	return nil
}
func (s *userServices) Login(u models.User) (dto.Login, error) {

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
func (s *userServices) UpdateUsername(u models.User, t dto.Token) error {

	u.UUID = t.UUID

	err := s.Database.UpdateUsername(u)
	if err != nil {
		return err
	}

	return nil
}

func (s *userServices) UpdatePassword(u models.User) error {

	err := s.Database.UpdatePassword(u)
	if err != nil {
		return err
	}

	return nil
}
