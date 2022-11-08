package controllers

import (
	"net/http"
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
	err := h.UserServices.SignUp(c)
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

	result, err := h.UserServices.Login(c)
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

	err := h.UserServices.UpdateUsername(c)
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

	err := h.UserServices.UpdatePassword(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
