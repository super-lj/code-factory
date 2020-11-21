package dao

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB gorm.DB

func InitDB() {
	dsn := "user=gorm password=gorm dbname=codefactoryci port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gorm Postgre open failed: %v", err)
	}
	DB = *db
}
