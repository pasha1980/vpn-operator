package service

import (
	"time"
	"vpn-operator/config"
	"vpn-operator/domain/infrastructure/apiError"
	"vpn-operator/domain/operator/model"
)

type CustomerService struct {
}

func (s *CustomerService) CreateClient(service string, serverId int) (*model.Client, error) {
	var instance model.Instance
	err := config.DB.First(&instance, serverId).Error
	if err != nil {
		return nil, apiError.NewNotFoundError("Server not found", nil)
	}

	if !instance.IsSupportService(service) {
		return nil, apiError.NewBadRequestError("Service not supported", nil)
	}

	rawClient := model.Client{
		Service:  service,
		ServerID: uint(serverId),
		Date:     time.Now(),
	}
	config.DB.Save(&rawClient)

	client, err := instance.GetClientConfiguration(&rawClient)
	if err != nil {
		return nil, err
	}

	err = client.SaveConfigToStorage()
	if err != nil {
		return nil, err
	}

	config.DB.Save(client)

	return client, nil
}

func (s *CustomerService) DeleteClient(clientId int) error {
	var client model.Client
	err := config.DB.First(&client, clientId).Error
	if err != nil {
		return apiError.NewNotFoundError("Client not found", nil)
	}

	err = client.DeleteConfigFromStorage()
	if err != nil {
		return err
	}

	var instance model.Instance
	err = config.DB.First(&instance, client.ServerID).Error
	if err != nil {
		return apiError.NewNotFoundError("Server not found", nil)
	}

	err = instance.RemoveClientConfiguration(&client)
	if err != nil {
		return err
	}

	config.DB.Delete(&client)

	return nil
}
