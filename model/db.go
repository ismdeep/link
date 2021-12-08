package model

import (
	"github.com/ismdeep/link/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // load mysql driver
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(config.Config.DB.Dialect, config.Config.DB.DSN)
	if err != nil {
		panic(err)
	}
	DB = db

	DB.AutoMigrate(&Link{})
}
