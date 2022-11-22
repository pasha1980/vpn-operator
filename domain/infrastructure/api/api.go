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
	clientRoutes(h)

	return h.Start(config.Config.HttpAddress)
}

func instanceRoutes(h *echo.Echo) {

	c := controller.InstanceController{
		Service: &service.InstanceService{},
	}

	h.POST("/manager/hook", c.Hook)
}

func clientRoutes(h *echo.Echo) {
	c := controller.CustomerController{
		Service: &service.CustomerService{},
	}

	customerApi := h.Group("", authMiddleware)

	clientApi := customerApi.Group("/client")
	clientApi.POST("/:service/:server_id", c.CreateClient)
	clientApi.DELETE("/:id", c.DeleteClient)
	clientApi.GET("/:id", c.GetClient)
}
