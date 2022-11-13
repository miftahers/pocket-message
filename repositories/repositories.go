package repositories

import (
	"pocket-message/dto"
	"pocket-message/models"

	"github.com/google/uuid"
)

type IDatabase interface {
	SaveNewUser(models.User) error
	Login(username string, password string) (models.User, error)
	GetUserByUsername(username string) (models.User, error)
	UpdateUsername(user_model models.User) error
	UpdatePassword(user_model models.User) error
	SaveNewPocketMessage(pocket_message_model models.PocketMessage) error
	SaveNewRandomID(random_id_model models.PocketMessageRandomID) error
	GetPocketMessageByRandomID(random_id string) (dto.PocketMessageWithRandomID, error)
	UpdateVisitCount(pocket_message_with_random_id dto.PocketMessageWithRandomID) error
	UpdatePocketMessage(pocket_message_model models.PocketMessage) error
	DeletePocketMessage(pocket_message_uuid uuid.UUID) error
	GetPocketMessageByUserUUID(pocket_message_uuid uuid.UUID) ([]dto.OwnedMessage, error)
}
