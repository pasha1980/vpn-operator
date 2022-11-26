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

func routeMiddleware(next echo.HandlerFunc) echo.HandlerFunc { // todo
	return func(c echo.Context) error {
		var status = 404

		currentUri := c.Request().RequestURI
		currentMethod := c.Request().Method

		for _, route := range c.Echo().Routes() {
			if route.Path == currentUri {
				if route.Method == currentMethod {
					status = 0
					break
				} else {
					status = 405
				}
			}
		}

		if status == 404 {
			return apiError.NewRouteNotFoundError(nil)
		} else if status == 405 {
			return apiError.NewMethodNotSupportedError(nil)
		}

		return next(c)
	}
}
