package controllers

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/middleware"
	"pocket-message/models"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type MockPocketMessageServices struct{}

func (s *MockPocketMessageServices) NewPocketMessage(c echo.Context) error {
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

	return nil
}
func (s *MockPocketMessageServices) GetPocketMessageByRandomID(c echo.Context) (dto.PocketMessageWithRandomID, error) {
	rid := c.Param("random_id")
	if rid == "" {
		return dto.PocketMessageWithRandomID{}, errors.New("error, random_id parameter can not be empty")
	}

	return dto.PocketMessageWithRandomID{Title: "Ini Test", Content: "Ini juga Test"}, nil
}
func (s *MockPocketMessageServices) UpdatePocketMessage(c echo.Context) error {
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

	return nil
}
func (s *MockPocketMessageServices) DeletePocketMessage(c echo.Context) error {
	_, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return err
	}
	return nil
}
func (s *MockPocketMessageServices) GetUserPocketMessage(c echo.Context) ([]dto.OwnedMessage, error) {
	_, err := middleware.DecodeJWT(c)
	if err != nil {
		return nil, err
	}
	return []dto.OwnedMessage{
		{
			Title:   "halo dunia",
			Content: "halo kamu",
			Visit:   1,
		},
	}, nil
}
