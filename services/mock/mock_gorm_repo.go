package services

import (
	"errors"
	"pocket-message/dto"
	"pocket-message/models"

	"github.com/google/uuid"
)

type MockGorm struct{}

// User
func (db *MockGorm) SaveNewUser(u models.User) error {
	if u.Username == "" {
		return errors.New("username exist")
	}
	return nil
}
func (db *MockGorm) Login(u models.User) (models.User, error) {
	if u.Username == "" {
		return models.User{}, errors.New("record not found")
	}
	return models.User{Username: "Udin", Password: "PasswordnyaUdin"}, nil
}
func (db *MockGorm) UpdateUsername(u models.User) error {
	if u.Username == "" {
		return errors.New("username has taken")
	}
	return nil
}
func (db *MockGorm) UpdatePassword(u models.User) error {
	if u.Password == "" {
		return errors.New("password should not be empty")
	}
	return nil
}

// PocketMessage
func (db *MockGorm) SaveNewPocketMessage(pm models.PocketMessage) error {
	if pm.Title == "" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) SaveNewRandomID(rid models.PocketMessageRandomID) error {
	if rid.RandomID == "" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) GetPocketMessageByRandomID(rid string) (dto.PocketMessageWithRandomID, error) {
	if rid == "" {
		return dto.PocketMessageWithRandomID{}, errors.New("database errror")
	}
	return dto.PocketMessageWithRandomID{Title: "udin bahagia"}, nil
}
func (db *MockGorm) UpdateVisitCount(rid dto.PocketMessageWithRandomID) error {
	if rid.UUID.String() == "" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) UpdatePocketMessage(newMsg models.PocketMessage) error {
	if newMsg.Title == "" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) DeletePocketMessage(msgID uuid.UUID) error {
	if msgID.String() == "" {
		return errors.New("record not found")
	}
	return nil
}
func (db *MockGorm) GetPocketMessageByUserUUID(uuid uuid.UUID) ([]dto.OwnedMessage, error) {
	if uuid.String() == "" {
		return nil, errors.New("database error")
	}
	return []dto.OwnedMessage{
		{
			RandomID: "superIdol",
			Title:    "Idol super",
			Content:  "dede gemesh",
		},
	}, nil
}
