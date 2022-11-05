package controllers

import (
	"net/http"
	"pocket-message/services"

	"github.com/labstack/echo/v4"
)

type PocketMessageHandler interface {
	NewPocketMessage(echo.Context) error
	GetPocketMessageByRandomID(echo.Context) error
	UpdatePocketMessage(echo.Context) error
	DeletePocketMessage(echo.Context) error
	GetOwnedPocketMessage(echo.Context) error
}
type pocketMessageHandler struct {
	services.PocketMessageServices
}

// TODO NewPocketMessage Unit Test
func (h *pocketMessageHandler) NewPocketMessage(c echo.Context) error {

	err := h.PocketMessageServices.NewPocketMessage(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "created",
	})
}

// TODO GetPocketMessageByRandomID Unit Test
func (h *pocketMessageHandler) GetPocketMessageByRandomID(c echo.Context) error {
	if c.Param("random_id") == "" {
		return c.JSON(http.StatusNoContent, echo.Map{
			"message": "error, url parameters can not be empty",
		})
	}

	result, err := h.PocketMessageServices.GetPocketMessageByRandomID(c)
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

// TODO UpdatePocketMessage Unit Test
func (h *pocketMessageHandler) UpdatePocketMessage(c echo.Context) error {
	err := h.PocketMessageServices.UpdatePocketMessage(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "updated",
	})
}

// TODO DeletePocketMessage Unit Test
func (h *pocketMessageHandler) DeletePocketMessage(c echo.Context) error {
	err := h.PocketMessageServices.DeletePocketMessage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "deleted",
	})
}

// TODO GetPocketMessageByUUID Unit Test
func (h *pocketMessageHandler) GetOwnedPocketMessage(c echo.Context) error {
	result, err := h.PocketMessageServices.GetUserPocketMessage(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "success",
		"data":    result,
	})
}
