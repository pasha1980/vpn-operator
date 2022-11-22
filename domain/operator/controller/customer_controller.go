package controller

import (
	"github.com/labstack/echo/v4"
	"strconv"
	"vpn-operator/domain/infrastructure/apiError"
	"vpn-operator/domain/operator/controller/response_dto"
	"vpn-operator/domain/operator/service"
)

type CustomerController struct {
	Service *service.CustomerService
}

func (c *CustomerController) CreateClient(ctx echo.Context) error {
	vpnService := ctx.Param("service")
	serverId, err := strconv.Atoi(ctx.Param("server_id"))
	if err != nil {
		return apiError.NewBadRequestError("Server ID is invalid", nil)
	}

	client, err := c.Service.CreateClient(vpnService, serverId)
	if err != nil {
		return err
	}

	responseDto := response_dto.Client{
		ID:       client.ID,
		FileName: *client.ConfigFileName,
		Config:   *client.Config,
	}

	return ctx.JSON(201, &responseDto)
}

func (c *CustomerController) DeleteClient(ctx echo.Context) error {
	clientId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return apiError.NewBadRequestError("Client ID is invalid", nil)
	}

	err = c.Service.DeleteClient(clientId)
	if err != nil {
		return err
	}

	return ctx.NoContent(204)
}

func (c *CustomerController) GetClient(ctx echo.Context) error {
	clientId, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return apiError.NewBadRequestError("Client ID is invalid", nil)
	}

	client, err := c.Service.GetClientConfig(clientId)
	if err != nil {
		return err
	}

	responseDto := response_dto.Client{
		ID:       client.ID,
		FileName: *client.ConfigFileName,
		Config:   *client.Config,
	}

	return ctx.JSON(200, &responseDto)
}
