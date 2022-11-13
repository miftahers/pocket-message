package users

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
)

func NewUserServices(db repositories.IDatabase) IUserServices {
	return &userServices{IDatabase: db}
}

type IUserServices interface {
	SignUp(user models.User) error
	Login(user models.User) (dto.Login, error)
	UpdateUsername(user models.User, token dto.Token) error
	UpdatePassword(user models.User) error
}

type userServices struct {
	repositories.IDatabase
}

func (s *userServices) SignUp(user models.User) error {

	_, err := s.IDatabase.GetUserByUsername(user.Username)
	if err != nil {
		if err.Error() == "record not found" {
			user.UUID = uuid.New()
			err = s.IDatabase.SaveNewUser(user)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return errors.New("username exist, try another username")
}
func (s *userServices) Login(user models.User) (dto.Login, error) {

	user, err := s.IDatabase.Login(user.Username, user.Password)
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
func (s *userServices) UpdateUsername(user models.User, t dto.Token) error {
	_, err := s.IDatabase.GetUserByUsername(user.Username)
	if err != nil {
		if err.Error() == "record not found" {
			user.UUID = t.UUID
			err = s.IDatabase.UpdateUsername(user)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return errors.New("username exist, try another username")
}
func (s *userServices) UpdatePassword(u models.User) error {

	err := s.IDatabase.UpdatePassword(u)
	if err != nil {
		return err
	}

	return nil
}
