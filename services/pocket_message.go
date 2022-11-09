package services

import (
	"pocket-message/dto"
	"pocket-message/helper"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
)

func NewPocketMessageServices(db repositories.Database) PocketMessageServices {
	return &pmServices{Database: db}
}

type PocketMessageServices interface {
	NewPocketMessage(msg models.PocketMessage, token dto.Token) error
	GetPocketMessageByRandomID(randomID string) (dto.PocketMessageWithRandomID, error)
	UpdatePocketMessage(pocketMessage models.PocketMessage) error
	DeletePocketMessage(id uuid.UUID) error
	GetUserPocketMessage(token dto.Token) ([]dto.OwnedMessage, error)
}

type pmServices struct {
	repositories.Database
}

func (s *pmServices) NewPocketMessage(pm models.PocketMessage, t dto.Token) error {

	pm.UUID = uuid.New()
	pm.UserUUID = t.UUID

	err := s.Database.SaveNewPocketMessage(pm)
	if err != nil {
		return err
	}

	var rid models.PocketMessageRandomID
	rid.PocketMessageUUID = pm.UUID
	rid.RandomID = helper.GenerateRandomString(8)

	err = s.Database.SaveNewRandomID(rid)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) GetPocketMessageByRandomID(rid string) (dto.PocketMessageWithRandomID, error) {

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
func (s *pmServices) UpdatePocketMessage(pm models.PocketMessage) error {

	err := s.Database.UpdatePocketMessage(pm)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) DeletePocketMessage(id uuid.UUID) error {

	err := s.Database.DeletePocketMessage(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) GetUserPocketMessage(t dto.Token) ([]dto.OwnedMessage, error) {

	result, err := s.Database.GetPocketMessageByUserUUID(t.UUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
