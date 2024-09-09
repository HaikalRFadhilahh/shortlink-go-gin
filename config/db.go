package config

import (
	"fmt"

	"github.com/HaikalRFadhilahh/shortlink-go-gin/helper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func CreateConnection() (*gorm.DB, error) {
	var (
		DB_HOST     = helper.GetEnv("DB_HOST", "127.0.0.1")
		DB_PORT     = helper.GetEnv("DB_PORT", "3306")
		DB_USERNAME = helper.GetEnv("DB_USERNAME", "root")
		DB_PASSWORD = helper.GetEnv("DB_PASSWORD", "")
		DB_NAME     = helper.GetEnv("DB_NAME", "")
	)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", DB_USERNAME, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Database Refused!")
	}

	return db, nil
}
