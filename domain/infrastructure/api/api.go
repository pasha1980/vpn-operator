package api

import (
	"github.com/labstack/echo/v4"
	"vpn-operator/config"
	"vpn-operator/domain/operator/controller"
	"vpn-operator/domain/operator/service"
)

func InitHttp() error {
	h := echo.New()
	h.HTTPErrorHandler = httpErrorHandler()

	instanceRoutes(h)
	userRoutes(h)

	return h.Start(config.Config.HttpAddress)
}

func instanceRoutes(h *echo.Echo) {

	c := controller.InstanceController{
		Service: service.InstanceService{},
	}

	h.POST("/manager/hook", c.Hook)
}

func userRoutes(h *echo.Echo) {

}
