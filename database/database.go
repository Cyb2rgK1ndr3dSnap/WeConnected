package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Cyb2rgK1ndr3dSnap/we-connected/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//var dbClient *gorm.DB

func ConnectToDB() (*gorm.DB, error) {
	var err error

	if err := godotenv.Load(); err != nil {
		fmt.Println("No se pudo cargar el archivo .env")
		return nil, nil
	}

	db_hostname := os.Getenv("DB_HOST")
	db_name := os.Getenv("DB_DB")
	db_user := os.Getenv("DB_USER")
	db_pass := os.Getenv("DB_PASSWORD")

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?tls=true&interpolateParams=true", db_user, db_pass, db_hostname, db_name)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect DB")
	}

	db.AutoMigrate(&models.Company{})
	db.AutoMigrate(&models.User{})

	return db, nil
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		// Manejo de error
		return
	}
	sqlDB.Close()
}
