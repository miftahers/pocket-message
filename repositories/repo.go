package repositories

import (
	"pocket-message/dto"
	"pocket-message/models"

	"github.com/google/uuid"
)

type Database interface {
	SaveNewUser(models.User) error
	Login(models.User) (models.User, error)
	UpdateUsername(models.User) error
	UpdatePassword(models.User) error
	SaveNewPocketMessage(models.PocketMessage) error
	SaveNewRandomID(models.PocketMessageRandomID) error
	GetPocketMessageByRandomID(rid string) (dto.PocketMessageWithRandomID, error)
	UpdateVisitCount(rid dto.PocketMessageWithRandomID) error
	UpdatePocketMessage(newMsg models.PocketMessage) error
	DeletePocketMessage(msgID uuid.UUID) error
	GetPocketMessageByUserUUID(uuid uuid.UUID) ([]dto.OwnedMessage, error)
}
