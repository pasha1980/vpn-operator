package model

import (
	"gorm.io/gorm"
	"time"
)

type Client struct {
	gorm.Model
	ServerID uint
	Service  string
	Date     time.Time
	KeyFile  string
}
