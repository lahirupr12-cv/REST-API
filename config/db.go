package config

import (
	"github.com/lahirupr12-cv/gin_restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// github.com/mattn/go-sqlite3
	dsn := "host=localhost user=postgres password=lahiru12 dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	checKError(err)

	db.AutoMigrate(&models.User{})
	DB = db

}

func checKError(err error) {
	if err != nil {
		panic(err)
	}
}
