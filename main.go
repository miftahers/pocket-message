package main

import (
	"pocket-message/configs"
	"pocket-message/configs/database"
	"pocket-message/routes"
)

func main() {
	db, err := database.ConnectDB()
	if err != nil {
		panic(err)
	}

	err = database.MigrateDB(db)
	if err != nil {
		panic(err)
	}

	e := routes.Init(db)
	err = e.Start(configs.APIPort)
	if err != nil {
		panic(err)
	}
}
