package pocket_messages

import (
	"pocket-message/dto"
	"pocket-message/helper"
	"pocket-message/models"
	"pocket-message/repositories"

	"github.com/google/uuid"
)

func NewPocketMessageServices(db repositories.IDatabase) IPocketMessageServices {
	return &pmServices{IDatabase: db}
}

type IPocketMessageServices interface {
	NewPocketMessage(msg models.PocketMessage, token dto.Token) error
	GetPocketMessageByRandomID(randomID string) (dto.MsgForPublic, error)
	UpdatePocketMessage(pocketMessage models.PocketMessage) error
	DeletePocketMessage(id uuid.UUID) error
	GetUserPocketMessage(token dto.Token) ([]dto.OwnedMessage, error)
}

type pmServices struct {
	repositories.IDatabase
}

func (s *pmServices) NewPocketMessage(pm models.PocketMessage, t dto.Token) error {

	// Create new Pocket Message UUID
	pm.UUID = uuid.New()
	// Set User UUID from token User UUID
	pm.UserUUID = t.UUID

	// Save to DB
	err := s.IDatabase.SaveNewPocketMessage(pm)
	if err != nil {
		return err
	}

	// Create Random ID to access pocket message
	var rid models.PocketMessageRandomID
	rid.PocketMessageUUID = pm.UUID
	rid.RandomID = helper.GenerateRandomString(8)

	// Save Random ID to DB
	err = s.IDatabase.SaveNewRandomID(rid)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) GetPocketMessageByRandomID(rid string) (dto.MsgForPublic, error) {

	// Get Pocket Message using Random ID
	v, err := s.IDatabase.GetPocketMessageByRandomID(rid)
	if err != nil {
		return dto.MsgForPublic{}, err
	}

	err = s.IDatabase.UpdateVisitCount(v)
	if err != nil {
		return dto.MsgForPublic{}, err
	}

	// Convert into public consumption data
	result := dto.MsgForPublic{
		Title:   v.Title,
		Content: v.Content,
	}

	return result, nil
}
func (s *pmServices) UpdatePocketMessage(pm models.PocketMessage) error {

	err := s.IDatabase.UpdatePocketMessage(pm)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) DeletePocketMessage(id uuid.UUID) error {

	err := s.IDatabase.DeletePocketMessage(id)
	if err != nil {
		return err
	}

	return nil
}
func (s *pmServices) GetUserPocketMessage(t dto.Token) ([]dto.OwnedMessage, error) {

	result, err := s.IDatabase.GetPocketMessageByUserUUID(t.UUID)
	if err != nil {
		return nil, err
	}

	return result, nil
}
