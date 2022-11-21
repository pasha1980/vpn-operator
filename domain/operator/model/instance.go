package model

import (
	"encoding/json"
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
	Region            *string
}

func (i Instance) Ping() bool {
	resp, err := http.Get(i.HttpUrl + "/ping")
	if err != nil {
		return false
	}

	if resp.StatusCode != 200 {
		return false
	}

	return true
}

func (i Instance) GetStatus() (availableServices []string, version string, err error) {
	type instanceServiceResponse struct {
		Service map[string]bool `json:"service"`
		Version string          `json:"version"`
	}

	client := &http.Client{}
	req, _ := http.NewRequest("GET", i.HttpUrl+"/api/status", nil)
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
