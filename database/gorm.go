package database

import (
	"fmt"
	"os"
	"pocket-message/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	// DB_Address = os.Getenv("DB_ADDRESS")
	// DB_Name = os.Getenv("DB_NAME")
	DB_Address = "localhost:3306"
	DB_Name    = "pocket_message"
)

func SetEnv(key, def string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		return def
	}
	return val
}

func ConnectDB() (*gorm.DB, error) {
	connectionString := fmt.Sprintf("root:root@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		DB_Address, DB_Name)

	return gorm.Open(mysql.Open(connectionString), &gorm.Config{})
}

func MigrateDB(db *gorm.DB) error {
	return db.AutoMigrate(
		models.User{},
		models.PocketMessage{},
		models.PocketMessageRandomID{},
	)
}
