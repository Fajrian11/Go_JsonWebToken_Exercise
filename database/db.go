package database

import (
	"Go_JsonWebToken_Exercise/models"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	dbHost     string
	dbUser     string
	dbPassword string
	dbPort     string
	dbName     string
	db         *gorm.DB
)

func StartDB() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error While Load .env")
	}

	dbHost = os.Getenv("DB_HOST")
	dbUser = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbPort = os.Getenv("DB_PORT")
	dbName = os.Getenv("DB_NAME")
	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s?charset=utf8&parseTime=True", dbUser, dbHost, dbPort, dbName)
	fmt.Println(dsn)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Printf("ERROR: Failed to connect to Database -> %v\n", err)
	}
	fmt.Println("Sukses Koneksi Ke database")

	db.AutoMigrate(&models.User{}, &models.Product{})
}

func GetDB() *gorm.DB {
	return db
}
