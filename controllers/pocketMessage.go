package controllers

import (
	"errors"
	"net/http"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/services"

	"github.com/google/uuid"
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

	var pm models.PocketMessage
	err := c.Bind(&pm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if pm.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("title should not be empty"),
		})
	}
	if pm.Content == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("content should not be empty"),
		})
	}

	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	err = h.PocketMessageServices.NewPocketMessage(pm, t)
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

	rid := c.Param("random_id")
	if rid == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("random_id parameter can not be empty"),
		})
	}

	result, err := h.PocketMessageServices.GetPocketMessageByRandomID(rid)
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

	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("uuid invalid"),
		})
	}

	var pm models.PocketMessage
	err = c.Bind(&pm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	if pm.Title == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("title should not be empty"),
		})
	}
	if pm.Content == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("content should not be empty"),
		})
	}

	pm.UUID = id

	err = h.PocketMessageServices.UpdatePocketMessage(pm)
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
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("uuid invalid"),
		})
	}

	err = h.PocketMessageServices.DeletePocketMessage(id)
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
	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	result, err := h.PocketMessageServices.GetUserPocketMessage(t)
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
