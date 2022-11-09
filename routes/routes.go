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

	api := e.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.POST("/signup", uHandler.SignUp)                                                              // host:port/api/v1/users/signup
	users.POST("/login", uHandler.Login)                                                                // host:port/api/v1/users/login
	users.PUT("/reset-password", uHandler.UpdatePassword)                                               // host:port/api/v1/users/reset-password
	users.PUT("/change-username", uHandler.UpdateUsername, middleware.JWT([]byte(configs.TokenSecret))) // host:port/api/v1/users/change-username

	pocketMessages := v1.Group("/pocket-messages")
	pocketMessages.POST("", pmHandler.NewPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))            // host:port/api/v1/pocket-messages
	pocketMessages.PUT("/:uuid", pmHandler.UpdatePocketMessage, middleware.JWT([]byte(configs.TokenSecret)))    // host:port/api/v1/pocket-messages/:uuid
	pocketMessages.DELETE("/:uuid", pmHandler.DeletePocketMessage, middleware.JWT([]byte(configs.TokenSecret))) // host:port/api/v1/pocket-messages/:uuid
	pocketMessages.GET("", pmHandler.GetOwnedPocketMessage, middleware.JWT([]byte(configs.TokenSecret)))        // host:port/api/v1/pocket-messages

	msg := pocketMessages.Group("/msg")
	msg.GET("/:random_id", pmHandler.GetPocketMessageByRandomID) // host:port/api/v1/pocket-messages/msg/:random_id

	return e
}
