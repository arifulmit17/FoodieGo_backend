package user

import (
	"net/http"

	"foodiego/internal/user/dto"

	"github.com/labstack/echo/v5"
)

type UserHandler struct {
	userService UserService
}

func NewUserHandler(userService UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// POST /api/auth/register
func (h *UserHandler) Register(c *echo.Context) error {

	var req dto.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
		})
	}

	res, err := h.userService.Register(req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, res)
}

// POST /api/auth/login
func (h *UserHandler) Login(c *echo.Context) error {

	var req dto.LoginRequest

	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]any{
			"message": "Invalid request body",
		})
	}

	res, err := h.userService.Login(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}

// GET /api/auth/me
func (h *UserHandler) Me(c *echo.Context) error {

	userID, ok := c.Get("userID").(string)
	if !ok {
		return c.JSON(http.StatusUnauthorized, map[string]any{
			"message": "Unauthorized",
		})
	}

	res, err := h.userService.GetProfile(userID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]any{
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, res)
}
