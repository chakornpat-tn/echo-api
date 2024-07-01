package main

import (
	"echo-api/database"
	"echo-api/routes"
	"net/http"

	"github.com/labstack/echo/v4"
)

func init() {
	database.InitDatabase()
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"success": true})
	})

	userGroup := e.Group("/user")
	routes.UserGroup(userGroup)

	e.Logger.Fatal(e.Start(":8080"))
}
