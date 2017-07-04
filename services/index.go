package services

import (
	"net/http"

	"github.com/labstack/echo"
)

//RegisterRoutes Register all api routes
func RegisterRoutes(api *echo.Group) {
	api.GET("/", apiIndex)
	api.GET("/health", healthCheck)
	api.GET("/events", getEvents)
	api.GET("/events/:id", getEvent)
	api.DELETE("/events/:id", deleteEvent)
	api.POST("/events", createEvent)
	api.PUT("/events/:id", updateEvent)
}

func apiIndex(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Carbon API",
		"events":  "/events",
	})
}

func healthCheck(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status":  200,
		"message": "Carbon API OK",
		"version": "0.0.1",
	})
}
