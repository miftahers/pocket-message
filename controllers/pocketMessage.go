package controllers

import (
	"net/http"
	"pocket-message/services"

	"github.com/labstack/echo/v4"
)

func NewPocketMessageHandler(service services.PocketMessageServices) PocketMessageHandler {
	return &pocketMessageHandler{
		PocketMessageServices: service,
	}
}

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
func (h *pocketMessageHandler) GetPocketMessageByRandomID(c echo.Context) error {

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
func (h *pocketMessageHandler) DeletePocketMessage(c echo.Context) error {
	err := h.PocketMessageServices.DeletePocketMessage(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "deleted",
	})
}
func (h *pocketMessageHandler) GetOwnedPocketMessage(c echo.Context) error {
	result, err := h.PocketMessageServices.GetUserPocketMessage(c)
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
