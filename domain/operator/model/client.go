package model

import (
	"encoding/base64"
	"gorm.io/gorm"
	"os"
	"time"
	"vpn-operator/config"
)

type Client struct {
	gorm.Model
	ServerID       uint
	Service        string
	Date           time.Time
	IsActive       bool
	ConfigFileName *string
	Config         *string `gorm:"-"`
}

func (c *Client) SaveConfigToStorage() error {
	file, err := os.Create(config.Config.StoragePath + *c.ConfigFileName)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := base64.StdEncoding.DecodeString(*c.Config)
	if err != nil {
		return err
	}
	_, err = file.WriteString(string(content))
	return err
}

func (c *Client) DeleteConfigFromStorage() error {
	return os.Remove(config.Config.StoragePath + *c.ConfigFileName)
}
