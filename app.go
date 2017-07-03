package main

import (
	"github.com/codeaustin/carbon/models"
	"github.com/codeaustin/carbon/services"
	"github.com/codeaustin/carbon/utils/config"
	"github.com/codeaustin/carbon/utils/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	// Ensure DB and tables created
	tables := []db.TableInfo{models.CreateCategoryTable(), models.CreateEventTable()}

	db.Init()
	db.CreateTables(tables)

	e := echo.New()
	e.Use(middleware.Logger())

	api := e.Group("/v1")
	services.RegisterRoutes(api)

	PORT := ":" + config.Config.Port
	e.Logger.Fatal(e.Start(PORT))
}
