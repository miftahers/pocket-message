package controllers

import "pocket-message/services"

func NewPocketMessageHandler(service services.PocketMessageServices) PocketMessageHandler {
	return &pocketMessageHandler{
		PocketMessageServices: service,
	}
}
func NewUserHandler(service services.UserServices) UserHandler {
	return &userHandler{
		UserServices: service,
	}
}
