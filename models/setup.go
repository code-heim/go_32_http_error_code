package models

import (
	"gin_blogs/app_config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(app_config.Cfg.DB), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database!")
	}

	DB = database
}

func DBMigrate() {
	DB.AutoMigrate(&Blog{}, &User{})
}
