package users

import (
	"errors"
	"net/http"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/services/users"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	users.IUserServices
}

func (h *UserHandler) SignUp(c echo.Context) error {
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

	err = h.IUserServices.SignUp(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, echo.Map{
		"message": "created",
	})
}

func (h *UserHandler) Login(c echo.Context) error {

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

	result, err := h.IUserServices.Login(u)
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

func (h *UserHandler) UpdateUsername(c echo.Context) error {
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

	err = h.IUserServices.UpdateUsername(u, t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}

func (h *UserHandler) UpdatePassword(c echo.Context) error {
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

	err = h.IUserServices.UpdatePassword(u)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
	})
}
