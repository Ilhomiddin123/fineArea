package db

import (
	"fineArea/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var conn *gorm.DB

func Connect() {
	// Подключение к PostgreSQL базе данных с помощью GORM
	dsn := "host=localhost user=postgres password=12345 dbname=fine_area_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Создание таблицы VehicleData, если она не существует
	err = db.AutoMigrate(&models.VehicleData{})
	if err != nil {
		panic("failed to migrate database")
	}

	conn = db
}

func GetConn() *gorm.DB {
	return conn
}
