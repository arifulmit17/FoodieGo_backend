package main

import (
	"fmt"
	"net/http"

	"foodiego/internal/config"
	"foodiego/internal/user"

	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
	config.ConnectDB()

	e.GET("/", func(c *echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"message": "Hello, World!"})
	})
	user.RegisterRoutes(e, config.DB)
	for _, r := range e.Router().Routes() {
		fmt.Printf("%-6s %s\n", r.Method, r.Path)
	}
	if err := e.Start(":8080"); err != nil {
		e.Logger.Error("failed to start server", "error", err)
	}
}
