package pocket_messages

import (
	"errors"
	"net/http"
	"pocket-message/middleware"
	"pocket-message/models"
	"pocket-message/services/pocket_messages"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type PocketMessageHandler struct {
	pocket_messages.IPocketMessageServices
}

func (h *PocketMessageHandler) NewPocketMessage(c echo.Context) error {

	var pm models.PocketMessage

	// retrieve data from request
	err := c.Bind(&pm)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	// Validasi data
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

	// Get obj that include User_UUID from JWT
	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	err = h.IPocketMessageServices.NewPocketMessage(pm, t)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusCreated, echo.Map{
		"message": "created",
	})
}
func (h *PocketMessageHandler) GetPocketMessageByRandomID(c echo.Context) error {

	// Get random id from path parameter
	rid := c.Param("random_id")
	if rid == "" {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("random_id parameter can not be empty"),
		})
	}

	result, err := h.IPocketMessageServices.GetPocketMessageByRandomID(rid)
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
func (h *PocketMessageHandler) UpdatePocketMessage(c echo.Context) error {

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

	err = h.IPocketMessageServices.UpdatePocketMessage(pm)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "updated",
	})
}
func (h *PocketMessageHandler) DeletePocketMessage(c echo.Context) error {
	id, err := uuid.Parse(c.Param("uuid"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": errors.New("uuid invalid"),
		})
	}

	err = h.IPocketMessageServices.DeletePocketMessage(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, echo.Map{
		"message": "deleted",
	})
}
func (h *PocketMessageHandler) GetOwnedPocketMessage(c echo.Context) error {
	t, err := middleware.DecodeJWT(c)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{
			"message": err.Error(),
		})
	}

	result, err := h.IPocketMessageServices.GetUserPocketMessage(t)
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
