package repositories

import (
	"pocket-message/dto"
	"pocket-message/models"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type GormSql struct {
	DB *gorm.DB
}

func NewGorm(db *gorm.DB) Database {
	return &GormSql{
		DB: db,
	}
}

// User
func (db GormSql) SaveNewUser(user models.User) error {
	result := db.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
func (db GormSql) Login(user models.User) (models.User, error) {
	err := db.DB.Where("username = ? AND password = ?",
		user.Username, user.Password).First(&user).Error
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
func (db GormSql) UpdateUsername(user models.User) error {
	err := db.DB.Model(&user).Where("uuid = ?", user.UUID).
		Update("username", user.Username).Error
	if err != nil {
		return err
	}
	return nil
}
func (db GormSql) UpdatePassword(user models.User) error {
	err := db.DB.Model(&user).Where("username = ?", user.Username).
		Update("password", user.Password).Error
	if err != nil {
		return err
	}

	return nil
}

// Pocket Message
func (db GormSql) SaveNewPocketMessage(pm models.PocketMessage) error {
	err := db.DB.Save(&pm).Error
	if err != nil {
		return err
	}
	return nil
}
func (db GormSql) SaveNewRandomID(rid models.PocketMessageRandomID) error {
	err := db.DB.Save(&rid).Error
	if err != nil {
		return err
	}
	return nil
}
func (db GormSql) GetPocketMessageByRandomID(rid string) (dto.PocketMessageWithRandomID, error) {
	var result dto.PocketMessageWithRandomID
	err := db.DB.Model(&models.PocketMessage{}).
		Select("pocket_messages.UUID, pocket_messages.title, pocket_messages.content,pocket_message_random_id.visit, pocket_message_random_id.random_id").
		Joins("LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid").
		Where("pocket_message_random_id.random_id = ?", rid).
		First(&result).Error
	if err != nil {
		return dto.PocketMessageWithRandomID{}, err
	}
	return result, nil
}
func (db GormSql) UpdateVisitCount(rid dto.PocketMessageWithRandomID) error {
	err := db.DB.Model(&models.PocketMessageRandomID{}).Where("pocket_message_uuid = ?", rid.UUID).Update("visit", rid.Visit+1).Error
	if err != nil {
		return err
	}

	return nil
}
func (db GormSql) UpdatePocketMessage(newMsg models.PocketMessage) error {
	err := db.DB.Model(&newMsg).Where("uuid = ?", newMsg.UUID).Updates(models.PocketMessage{
		Title:   newMsg.Title,
		Content: newMsg.Content,
	}).Error
	if err != nil {
		return err
	}
	return nil
}
func (db GormSql) DeletePocketMessage(msgID uuid.UUID) error {
	err := db.DB.Unscoped().Delete(&models.PocketMessage{}, "uuid = ?", msgID).Error
	if err != nil {
		return err
	}
	return nil
}
func (db GormSql) GetPocketMessageByUserUUID(uuid uuid.UUID) ([]dto.OwnedMessage, error) {
	var result []dto.OwnedMessage
	err := db.DB.Model(&models.PocketMessage{}).
		Select("pocket_message_random_id.random_id, pocket_messages.title, pocket_messages.content, pocket_message_random_id.visit").
		Joins("LEFT JOIN pocket_message_random_id ON pocket_messages.uuid = pocket_message_random_id.pocket_message_uuid").
		Where("pocket_messages.user_uuid = ?", uuid).
		Scan(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
