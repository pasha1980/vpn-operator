package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func initDatabase() {
	db, err := gorm.Open(mysql.Open(Config.DatabaseLink), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
