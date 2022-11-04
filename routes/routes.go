package routes

import (
	"pocket-message/configs"
	"pocket-message/controllers"
	mid "pocket-message/middleware"
	"pocket-message/repositories"
	"pocket-message/services"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
)

func Init(db *gorm.DB) *echo.Echo {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	mid.LogMiddleware(e)

	repo := repositories.NewGorm(db)
	userServ, pmServ := services.NewServices(repo)
	uHandler := controllers.NewUserHandler(userServ)
	pmHandler := controllers.NewPocketMessageHandler(pmServ)

	e.POST("/signup", uHandler.SignUp)
	e.POST("/login", uHandler.Login)
	e.PUT("/users/reset-password", uHandler.UpdatePassword)
	e.PUT("/users/change-username", uHandler.UpdateUsername, middleware.JWT([]byte(configs.TokenSecret)))
	e.POST("/pocket-messages", pmHandler.NewPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))
	e.GET("/msg/:random_id", pmHandler.GetPocketMessageByRandomID)
	e.PUT("/pocket-messages/:uuid", pmHandler.UpdatePocketMessage, middleware.JWT([]byte(configs.TokenSecret)))
	e.DELETE("/pocket-messages/:uuid", pmHandler.DeletePocketMessage, middleware.JWT([]byte(configs.TokenSecret)))
	e.GET("pocket-messages", pmHandler.GetOwnedPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))

	return e
}
