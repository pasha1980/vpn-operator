package api

import (
	"github.com/labstack/echo/v4"
	"vpn-operator/domain/infrastructure/apiError"
	"vpn-operator/domain/infrastructure/auth"
)

func authMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		token := c.Request().Header.Get("Authorization")

		isValid, err := auth.CheckApiToken(token)
		if !isValid || err != nil {
			return apiError.NewAccessDeniedError("Invalid token")
		}

		return next(c)
	}
}
