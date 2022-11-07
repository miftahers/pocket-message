package services

import (
	"pocket-message/repositories"
)

func NewServices(db repositories.Database) (UserServices, PocketMessageServices) {
	return &userServices{Database: db},
		&pmServices{Database: db}
}
