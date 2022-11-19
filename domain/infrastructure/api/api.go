package api

import (
	"github.com/labstack/echo/v4"
	"vpn-operator/config"
)

func InitHttp() error {
	h := echo.New()
	h.HTTPErrorHandler = httpErrorHandler()

	// routes

	return h.Start(config.Config.HttpAddress)
}
