package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v5"
)

func JWTMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c *echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")

		if authHeader == "" {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "missing authorization header",
			})
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if tokenString == authHeader {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid authorization header",
			})
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid token",
			})
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "invalid claims",
			})
		}

		// Save user ID into Echo context
		c.Set("userID", claims["user_id"])

		return next(c)
	}
}
