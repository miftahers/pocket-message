package main

import (
	"pocket-message/configs"
	configDB "pocket-message/configs/database"
	"pocket-message/routes"
	v1 "pocket-message/routes/v1"
)

func main() {

	configs.InitConfig()
	configDB.InitDatabase()

	routePayload := &routes.Payload{
		DBGorm: configDB.DB,
		Config: configs.Cfg,
	}

	routePayload.InitUserService()

	e, trace := v1.InitRoute(routePayload)
	defer trace.Close()

	err := e.Start(configs.Cfg.APIPort)
	if err != nil {
		panic(err)
	}
}
