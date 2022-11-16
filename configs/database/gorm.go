package database

import (
	"fmt"
	"log"
	"pocket-message/configs"
	"pocket-message/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {

	DB_User := configs.Cfg.DB_USER
	DB_Password := configs.Cfg.DB_PASSWORD
	DB_Host := configs.Cfg.DB_HOST
	DB_Port := configs.Cfg.DB_PORT
	DB_Name := configs.Cfg.DB_NAME

	connectionString := fmt.
		Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
			DB_User,
			DB_Password,
			DB_Host,
			DB_Port,
			DB_Name)

	db, err := gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	DB = db

	err = DB.AutoMigrate(
		&models.User{},
		&models.PocketMessage{},
		&models.PocketMessageRandomID{},
	)
	if err != nil {
		panic(err)
	}

	log.Print("Init DB Done")
}
