package model

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"time"
)

type Instance struct {
	gorm.Model
	IP                string `gorm:"unique"`
	HttpUrl           string
	Version           string
	IsActive          bool
	UpDate            *time.Time
	Secret            string
	AvailableServices []string `gorm:"json"`
	Country           string
	Region            string
	City              string
}

func (i *Instance) Ping() bool {
	resp, err := http.Get(i.HttpUrl + "/ping")
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func (i *Instance) GetStatus() (availableServices []string, version string, err error) {
	type instanceServiceResponse struct {
		Service map[string]bool `json:"service"`
		Version string          `json:"version"`
	}

	client := &http.Client{}
	req, _ := http.NewRequest(http.MethodGet, i.HttpUrl+"/api/status", nil)
	req.Header.Set("Authorization", i.Secret)
	resp, err := client.Do(req)
	if err != nil {
		return nil, "", err
	}
	defer resp.Body.Close()

	var instanceStatus instanceServiceResponse
	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &instanceStatus)

	i.Version = instanceStatus.Version
	i.AvailableServices = []string{}
	for service, status := range instanceStatus.Service {
		if status {
			i.AvailableServices = append(i.AvailableServices, service)
		}
	}

	return i.AvailableServices, i.Version, nil
}

func (i *Instance) IsSupportService(service string) bool {
	for _, availableService := range i.AvailableServices {
		if availableService == service {
			return true
		}
	}
	return false
}

func (i *Instance) GetClientConfiguration(client *Client) (*Client, error) {
	type instanceCreateClientResponse struct {
		ID       int    `json:"id"`
		FileName string `json:"fileName"`
		Config   string `json:"client"`
	}
	var clientData instanceCreateClientResponse

	clientIdParam := fmt.Sprintf("%d", client.ID)

	httpClient := &http.Client{}
	req, _ := http.NewRequest(http.MethodPost, i.HttpUrl+"/api/"+client.Service+"/client/"+clientIdParam, nil)
	req.Header.Set("Authorization", i.Secret)
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)
	json.Unmarshal(respBody, &clientData)

	client.Config = &clientData.Config
	client.ConfigFileName = &clientData.FileName

	return client, nil
}

func (i *Instance) RemoveClientConfiguration(client *Client) error {
	clientIdParam := fmt.Sprintf("%d", client.ID)
	httpClient := &http.Client{}
	req, _ := http.NewRequest(http.MethodDelete, i.HttpUrl+"/api/"+client.Service+"/client/"+clientIdParam, nil)
	req.Header.Set("Authorization", i.Secret)
	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return nil
}
