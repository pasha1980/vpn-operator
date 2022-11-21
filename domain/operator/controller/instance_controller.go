package controller

import (
	"github.com/labstack/echo/v4"
	"vpn-operator/domain/infrastructure/apiError"
	"vpn-operator/domain/operator/controller/request_dto"
	"vpn-operator/domain/operator/service"
)

type InstanceController struct {
	Service *service.InstanceService
}

func (c InstanceController) Hook(ctx echo.Context) error {
	var hook request_dto.InstanceHook
	err := ctx.Bind(&hook)
	if err != nil {
		return err
	}

	var hookError *apiError.InstanceHookError
	switch hook.Action {
	case "up":
		hookError = c.Service.InstanceUp(
			ctx.RealIP(),
			hook.URL,
			hook.AvailableServices,
			*hook.Secret,
			*hook.Version,
			*hook.Country,
			*hook.Region,
			*hook.City,
		)
		break

	case "down":
		hookError = c.Service.InstanceDown(
			ctx.RealIP(),
			hook.URL,
		)
		break
	}

	if hookError != nil {
		return hookError
	}

	return ctx.NoContent(200)
}
