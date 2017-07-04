package services

import (
	"net/http"
	"strconv"

	"github.com/codeaustin/carbon/models"
	"github.com/labstack/echo"
)

//getEvents returns all events with a limit and offset
func getEvents(c echo.Context) error {
	limit, err := strconv.ParseUint(c.QueryParam("limit"), 10, 32)
	if err != nil {
		limit = 25
	}

	offset, err := strconv.ParseUint(c.QueryParam("offset"), 10, 32)
	if err != nil {
		offset = 0
	}

	offset *= limit

	events, tx := models.GetEvents(limit, offset)

	if !tx.Ok {
		return c.JSON(tx.Status, map[string]interface{}{
			"mesage": tx.Message,
		})
	}

	return c.JSON(tx.Status, map[string]interface{}{
		"events": events,
		"limit":  limit,
		"offset": offset,
	})
}

func getEvent(c echo.Context) error {
	id := c.Param("id")
	event, tx := models.GetEvent(id)

	if !tx.Ok {
		return c.JSON(tx.Status, map[string]interface{}{
			"mesage": tx.Message,
		})
	}

	return c.JSON(tx.Status, map[string]interface{}{
		"event": event,
	})
}

func deleteEvent(c echo.Context) error {
	id := c.Param("id")
	tx := models.DeleteEvent(id)

	return c.JSON(tx.Status, map[string]interface{}{
		"message": tx.Message,
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

func updateEvent(c echo.Context) error {
		id := c.Param("id")

		var fieldsMap map[string]interface{}
		if err := c.Bind(&fieldsMap); err != nil {
			return c.JSON(http.StatusBadRequest, map[string]interface{}{
				"message": err.Error(),
			})
		}

		tx := models.UpdateEvent(id, fieldsMap)

		if !tx.Ok {
			return c.JSON(tx.Status, map[string]interface{}{
				"message": tx.Message,
			})
		}

		return c.JSON(tx.Status, map[string]interface{}{
			"message": tx.Message,
			"updated fields": fieldsMap,
		})
}
