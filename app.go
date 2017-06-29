package main

import (
	"net/http"
	
	"github.com/codeaustin/carbon/utils/config"
	"github.com/codeaustin/carbon/utils/db"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	config.Load("/Users/cody/goworkspace/src/github.com/codeaustin/carbon/config")
	db.Init()

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	PORT := ":" + config.Config.Port
	e.Logger.Fatal(e.Start(PORT))
}