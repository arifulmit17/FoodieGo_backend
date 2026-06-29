package user

import (
	"fmt"

	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

func RegisterRoutes(e *echo.Echo, db *gorm.DB) {
	fmt.Println(db)
	userRepository := NewUserRepository(db)
	userService := NewUserService(userRepository)
	userHandler := NewUserHandler(userService)
	api := e.Group("/api/v1/auth")

	api.POST("/register", userHandler.Register)
	api.POST("/login", userHandler.Login)
	api.GET("/me", userHandler.Me)
}
