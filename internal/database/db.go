package database

import (
	"fmt"
	"os"

	"MarcoZillgen/homeChef/internal/storage"

	"joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() (*gorm.DB, error) {
	
	iferr := godotenv.Load("../../.env");err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// dsn := "host=localhost user=server password=password dbname=home_chef port=5432 sslmode=disable"
	dsn := fmt.Sprintf("host=localhost user=%s password=%s dbname=home_chef port=5432 sslmode=disable", user, password)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	if err = db.AutoMigrate(&storage.StorageItem{}); err != nil {
		return nil, fmt.Errorf("failed to migrate database: %v", err)
	}

	fmt.Println("Successfully connected to the database and migrated schema!")
	return db, nil
}