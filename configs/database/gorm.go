package database

import (
	"fmt"
	"os"
	"pocket-message/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB_User     = os.Getenv("DB_USER")
	DB_Password = os.Getenv("DB_PASSWORD")
	DB_Address  = os.Getenv("DB_ADDRESS")
	DB_Host     = os.Getenv("DB_HOST")
	DB_Port     = os.Getenv("DB_PORT")
	DB_Name     = os.Getenv("DB_NAME")
)

/*
	func SetEnv(key, def string) string {
		val, ok := os.LookupEnv(key)
		if !ok {
			return def
		}
		return val
	}
*/

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_User, DB_Password, DB_Host, DB_Port, DB_Name)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.PocketMessage{},
		models.PocketMessageRandomID{},
	)
}
