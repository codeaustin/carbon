package services

import (
	"net/http"

	"github.com/codeaustin/carbon/models"
	"github.com/labstack/echo"
)

func getEvents(c echo.Context) error {
	events, tx := models.GetEvents()

	if !tx.Ok {
		return c.JSON(tx.Status, map[string]interface{}{
			"mesage": tx.Message,
		})
	}

	return c.JSON(tx.Status, map[string]interface{}{
		"events": events,
	})
}

func createEvent(c echo.Context) error {
	var event models.Event

	if err := c.Bind(&event); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
		})
	}

	tx := models.CreateEvent(&event)

	if !tx.Ok {
		return c.JSON(tx.Status, map[string]interface{}{
			"message": tx.Message,
		})
	}

	return c.JSON(tx.Status, map[string]interface{}{
		"message": tx.Message,
		"event":   event,
	})
}