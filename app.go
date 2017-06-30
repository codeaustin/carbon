package main

import (
	"net/http"

	"github.com/codeaustin/carbon/models"
	"github.com/codeaustin/carbon/utils/config"
	"github.com/codeaustin/carbon/utils/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Init the config instance
	config.Init()

	// Ensure DB and tables created
	tables := []db.TableInfo{models.CreateCategoryTable()}

	db.Init()
	db.CreateTables(tables)

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hello World!",
		})
	})

	PORT := ":" + config.Config.Port
	e.Logger.Fatal(e.Start(PORT))
}
