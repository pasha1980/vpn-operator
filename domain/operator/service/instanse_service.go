package service

import (
	"time"
	"vpn-operator/config"
	"vpn-operator/domain/infrastructure/apiError"
	"vpn-operator/domain/operator/model"
)

type InstanceService struct {
}

func (s InstanceService) InstanceUp(
	IP string,
	HttpURL string,
	AvailableServices []string,
	Secret string,
	Version string,
	Country string,
	Region string,
	City string,
) *apiError.InstanceHookError {
	var instance model.Instance

	var existingInstance model.Instance
	config.DB.Where(&model.Instance{
		IP: IP,
	}).First(&existingInstance)

	if existingInstance.ID != 0 {
		instance = existingInstance
		instance.AvailableServices = AvailableServices
		instance.Secret = Secret
		instance.Version = Version
		instance.Country = Country
		instance.Region = Region
		instance.City = City
		instance.HttpUrl = HttpURL
	} else {
		instance = model.Instance{
			IP:                IP,
			HttpUrl:           HttpURL,
			AvailableServices: AvailableServices,
			Secret:            Secret,
			Version:           Version,
			Country:           Country,
			Region:            Region,
			City:              City,
		}

		if instance.Ping() {
			instance.IsActive = true
			currentDate := time.Now()
			instance.UpDate = &currentDate
		} else {
			instance.IsActive = false
			instance.UpDate = nil
		}
	}

	config.DB.Save(&instance)
	return nil
}

func (s InstanceService) InstanceDown(
	IP string,
	HttpUrl string,
) *apiError.InstanceHookError {
	var instance model.Instance
	config.DB.Where(&model.Instance{
		IP: IP,
	}).First(&instance)

	if instance.ID == 0 {
		return nil
	}

	instance.IsActive = false
	config.DB.Save(&instance)
	return nil
}
