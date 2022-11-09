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
	if u.Username == "admin" {
		return errors.New("username has been taken")
	}
	return nil
}
func (db *MockGorm) Login(u models.User) (models.User, error) {
	if u.Username == "suneo" {
		return models.User{}, errors.New("record not found")
	}
	return models.User{
		UUID:     uuid.Nil,
		Username: u.Username,
		Password: u.Password,
	}, nil
}
func (db *MockGorm) UpdateUsername(u models.User) error {
	if u.Username == "suneo" {
		return errors.New("username has taken")
	}
	return nil
}
func (db *MockGorm) UpdatePassword(u models.User) error {
	if len(u.Password) < 8 {
		return errors.New("password too short")
	}
	return nil
}

// PocketMessage
func (db *MockGorm) SaveNewPocketMessage(pm models.PocketMessage) error {
	if pm.Title == "super" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) SaveNewRandomID(rid models.PocketMessageRandomID) error {
	if len(rid.RandomID) == 16 {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) GetPocketMessageByRandomID(rid string) (dto.PocketMessageWithRandomID, error) {
	if rid == "superidol" {
		return dto.PocketMessageWithRandomID{}, errors.New("record not found")
	} else if rid == "igantenk" {
		return dto.PocketMessageWithRandomID{}, nil
	}
	return dto.PocketMessageWithRandomID{
		Title: "selsya bahagia",
	}, nil
}
func (db *MockGorm) UpdateVisitCount(rid dto.PocketMessageWithRandomID) error {
	if rid.UUID.String() == uuid.Nil.String() {
		return errors.New("record not found")
	}
	return nil
}
func (db *MockGorm) UpdatePocketMessage(newMsg models.PocketMessage) error {
	if newMsg.Title == "super" {
		return errors.New("database error")
	}
	return nil
}
func (db *MockGorm) DeletePocketMessage(msgID uuid.UUID) error {
	if msgID.String() == uuid.Nil.String() {
		return errors.New("record not found")
	}
	return nil
}
func (db *MockGorm) GetPocketMessageByUserUUID(id uuid.UUID) ([]dto.OwnedMessage, error) {
	if id == uuid.Nil {
		return nil, errors.New("record not found")
	}
	return []dto.OwnedMessage{
		{
			RandomID: "akasupas",
			Title:    "vtuber",
			Content:  "donation",
			Visit:    1000,
		},
	}, nil
}
