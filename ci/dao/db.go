package dao

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var RunDB *gorm.DB

func init() {
	dsn := "host=localhost port=5432 user=cfci password=cfci dbname=codefactoryci sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Gorm Postgre open failed: %v", err)
	}
	RunDB = db
}
