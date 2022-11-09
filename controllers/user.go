package controllers

import (
	"errors"
	"net/http"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/services"

	"github.com/labstack/echo/v4"
)

func NewUserHandler(service services.UserServices) UserHandler {
	return &userHandler{
		UserServices: service,
	}
}

type UserHandler interface {
	SignUp(echo.Context) error
	Login(echo.Context) error
	UpdateUsername(echo.Context) error
	UpdatePassword(echo.Context) error
}

type userHandler struct {
	services.UserServices
}

func (h *userHandler) SignUp(c echo.Context) error {
	// validation
	var u models.User

	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if u.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("username should not be empty"),
		})
	}
	if u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("password should not be empty"),
		})
	}

	err = h.UserServices.SignUp(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "created",
	})
}

func (h *userHandler) Login(c echo.Context) error {

	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if u.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("username should not be empty"),
		})
	}
	if u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("password should not be empty"),
		})
	}

	result, err := h.UserServices.Login(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    result,
	})
}

func (h *userHandler) UpdateUsername(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if u.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("error, username should not be empty"),
		})
	}

	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = h.UserServices.UpdateUsername(u, t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (h *userHandler) UpdatePassword(c echo.Context) error {
	var u models.User
	err := c.Bind(&u)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if u.Username == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("username should not be empty"),
		})
	}
	if u.Password == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("password should not be empty"),
		})
	}

	err = h.UserServices.UpdatePassword(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
