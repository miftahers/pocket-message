package repositories

import "pocket-message/models"

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
