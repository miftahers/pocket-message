package controllers

import (
	"net/http"
	"pocket-message/services"

	"github.com/labstack/echo/v4"
)

type UserHandler interface {
	SignUp(echo.Context) error
	Login(echo.Context) error
	UpdateUsername(echo.Context) error
	UpdatePassword(echo.Context) error
}

type userHandler struct {
	services.UserServices
}

// TODO SignUp Unit Test
func (h *userHandler) SignUp(c echo.Context) error {
	// validation
	err := h.UserServices.SignUp(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "success",
	})
}

// TODO Login Unit Test
func (h *userHandler) Login(c echo.Context) error {

	result, err := h.UserServices.Login(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    result,
	})
}

// TODO UpdateUsername Unit Test
func (h *userHandler) UpdateUsername(c echo.Context) error {

	err := h.UserServices.UpdateUsername(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

// TODO UpdatePassword Unit Test
func (h *userHandler) UpdatePassword(c echo.Context) error {

	err := h.UserServices.UpdatePassword(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": "error",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
