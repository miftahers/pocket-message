package routes

import (
	"database/sql"
	"pocket-message/configs"
	"pocket-message/repositories"
	pmService "pocket-message/services/pocket_messages"
	uService "pocket-message/services/users"

	"gorm.io/gorm"
)

type Payload struct {
	Config    *configs.Config
	DBGorm    *gorm.DB
	DBSql     *sql.DB
	repoSql   repositories.IDatabase
	uService  uService.IUserServices
	pmService pmService.IPocketMessageServices
}

func (p *Payload) InitUserService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.uService = uService.NewUserServices(p.repoSql)
}
func (p *Payload) InitPocketMessageService() {
	if p.repoSql == nil {
		p.InitRepoMysql()
	}

	p.uService = uService.NewUserServices(p.repoSql)
}

func (p *Payload) InitRepoMysql() {
	p.repoSql = repositories.NewGorm(p.DBGorm)
}

func (p *Payload) GetUserServices() uService.IUserServices {
	if p.uService == nil {
		p.InitUserService()
	}
	return p.uService
}

func (p *Payload) GetPocketMessageServices() pmService.IPocketMessageServices {
	if p.pmService == nil {
		p.InitPocketMessageService()
	}
	return p.pmService
}
