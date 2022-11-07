package services

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/helper"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PocketMessageServices interface {
	NewPocketMessage(echo.Context) error
	GetPocketMessageByRandomID(echo.Context) (dto.PocketMessageWithRandomID, error)
	UpdatePocketMessage(echo.Context) error
	DeletePocketMessage(echo.Context) error
	GetUserPocketMessage(echo.Context) ([]dto.OwnedMessage, error)
}

type pmServices struct {
	repositories.Database
}

// TODO NewPocketMessage Unit Test
func (s *pmServices) NewPocketMessage(c echo.Context) error {

	var pm models.PocketMessage
	err := c.Bind(&pm)
	if err != nil {
		return err
	}

	if pm.Title == "" {
		return errors.New("error, title should not be empty")
	}
	if pm.Content == "" {
		return errors.New("error, content should not be empty")
	}

	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return err
	}
	if c.Request().Header.Get("test") == "true" {
		pm.UUID = uuid.Nil
	} else {
		pm.UUID = uuid.New()
	}

	pm.UserUUID = t.UUID

	err = s.Database.SaveNewPocketMessage(pm)
	if err != nil {
		return err
	}

	var rid models.PocketMessageRandomID
	rid.PocketMessageUUID = pm.UUID
	if c.Request().Header.Get("test") == "true" {
		rid.RandomID = "test"
	} else {
		rid.RandomID = helper.GenerateRandomString(8)
	}

	err = s.Database.SaveNewRandomID(rid)
	if err != nil {
		return err
	}

	return nil
}

// TODO GetPocketMessageByRandomID Unit Test
func (s *pmServices) GetPocketMessageByRandomID(c echo.Context) (dto.PocketMessageWithRandomID, error) {

	rid := c.Param("random_id")
	if rid == "" {
		return dto.PocketMessageWithRandomID{}, errors.New("error, random_id parameter can not be empty")
	}

	result, err := s.Database.GetPocketMessageByRandomID(rid)
	if err != nil {
		return dto.PocketMessageWithRandomID{}, err
	}

	err = s.Database.UpdateVisitCount(result)
	if err != nil {
		return dto.PocketMessageWithRandomID{}, err
	}

	return result, nil
}

// TODO UpdatePocketMessage Unit Test
func (s *pmServices) UpdatePocketMessage(c echo.Context) error {
	var pm models.PocketMessage
	err := c.Bind(&pm)
	if err != nil {
		return err
	}

	if pm.Title == "" {
		return errors.New("error, title should not be empty")
	}
	if pm.Content == "" {
		return errors.New("error, content should not be empty")
	}

	pm.UUID, err = uuid.Parse(c.Param("uuid"))
	if err != nil {
		return err
	}

	err = s.Database.UpdatePocketMessage(pm)
	if err != nil {
		return err
	}

	return nil
}

// TODO DeletePocketMessage Unit Test
func (s *pmServices) DeletePocketMessage(c echo.Context) error {
	uuid, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return err
	}
	err = s.Database.DeletePocketMessage(uuid)
	if err != nil {
		return err
	}

	return nil
}

// TODO GetUserPocketMessage Unit Test
func (s *pmServices) GetUserPocketMessage(c echo.Context) ([]dto.OwnedMessage, error) {
	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return nil, err
	}
	result, err := s.Database.GetPocketMessageByUserUUID(t.UUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
