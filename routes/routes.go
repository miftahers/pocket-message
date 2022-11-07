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
	userServ := services.NewUserServices(repo)
	pmServ := services.NewPocketMessageServices(repo)
	uHandler := controllers.NewUserHandler(userServ)
	pmHandler := controllers.NewPocketMessageHandler(pmServ)

	api := e.Group("/api")                                                                                          // host:port/api/...
	v1 := api.Group("/v1")                                                                                          // host:port/api/v1/...
	v1.POST("/signup", uHandler.SignUp)                                                                             // host:port/api/v1/signup
	v1.POST("/login", uHandler.Login)                                                                               // host:port/api/v1/login
	v1.PUT("/users/reset-password", uHandler.UpdatePassword)                                                        // host:port/api/v1/users/reset-password
	v1.PUT("/users/change-username", uHandler.UpdateUsername, middleware.JWT([]byte(configs.TokenSecret)))          // host:port/api/v1/users/change-username
	v1.POST("/pocket-messages", pmHandler.NewPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))            // host:port/api/v1/pocket-messages
	v1.GET("/msg/:random_id", pmHandler.GetPocketMessageByRandomID)                                                 // host:port/api/v1/msg/:random_id
	v1.PUT("/pocket-messages/:uuid", pmHandler.UpdatePocketMessage, middleware.JWT([]byte(configs.TokenSecret)))    // host:port/api/v1/pocket-messages/:uuid
	v1.DELETE("/pocket-messages/:uuid", pmHandler.DeletePocketMessage, middleware.JWT([]byte(configs.TokenSecret))) // host:port/api/v1/pocket-messages/:uuid
	v1.GET("/pocket-messages", pmHandler.GetOwnedPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))        // host:port/api/v1/pocket-messages

	return e
}
