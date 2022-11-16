package v1

import (
	"io"
	"pocket-message/controllers/rest-echo/pocket_messages"
	"pocket-message/controllers/rest-echo/users"
	mid "pocket-message/middleware"
	"pocket-message/routes"

	"github.com/labstack/echo-contrib/jaegertracing"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func InitRoute(payload *routes.Payload) (*echo.Echo, io.Closer) {
	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	mid.LogMiddleware(e)
	trace := jaegertracing.New(e, nil)

	uHandler := users.UserHandler{
		IUserServices: payload.GetUserServices(),
	}
	pmHandler := pocket_messages.PocketMessageHandler{
		IPocketMessageServices: payload.GetPocketMessageServices(),
	}

	api := e.Group("/api")
	v1 := api.Group("/v1")

	users := v1.Group("/users")
	users.POST("/signup", uHandler.SignUp)                                                                     // host:port/api/v1/users/signup
	users.POST("/login", uHandler.Login)                                                                       // host:port/api/v1/users/login
	users.PUT("/reset-password", uHandler.UpdatePassword)                                                      // host:port/api/v1/users/reset-password
	users.PUT("/change-username", uHandler.UpdateUsername, middleware.JWT([]byte(payload.Config.TokenSecret))) // host:port/api/v1/users/change-username

	pocketMessages := v1.Group("/pocket-messages")
	pocketMessages.POST("", pmHandler.NewPocketMessage, middleware.JWT([]byte(payload.Config.TokenSecret)))            // host:port/api/v1/pocket-messages
	pocketMessages.PUT("/:uuid", pmHandler.UpdatePocketMessage, middleware.JWT([]byte(payload.Config.TokenSecret)))    // host:port/api/v1/pocket-messages/:uuid
	pocketMessages.DELETE("/:uuid", pmHandler.DeletePocketMessage, middleware.JWT([]byte(payload.Config.TokenSecret))) // host:port/api/v1/pocket-messages/:uuid
	pocketMessages.GET("", pmHandler.GetOwnedPocketMessage, middleware.JWT([]byte(payload.Config.TokenSecret)))        // host:port/api/v1/pocket-messages

	msg := pocketMessages.Group("/msg")
	msg.GET("/:random_id", pmHandler.GetPocketMessageByRandomID) // host:port/api/v1/pocket-messages/msg/:random_id

	return e, trace
}
