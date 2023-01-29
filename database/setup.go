package database

import (
	"fmt"
	"log"
	"os"

	"github.com/Shiroyasha19/task-5-vix-btpns-AdjiPrayoga/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

// Melakukan set up database berdasarkan env dan menyambungkan dengan db postgre sql
func SetupDB() *gorm.DB {

	godotenv.Load(".env")

	DB_HOST := os.Getenv("DB_HOST")
	DB_DRIVER := os.Getenv("DB_DRIVER")
	DB_USER := os.Getenv("DB_USER")
	DB_PASSWORD := os.Getenv("DB_PASSWORD")
	DB_NAME := os.Getenv("DB_NAME")
	DB_PORT := os.Getenv("DB_PORT")

	// URL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DB_HOST, DB_PORT, DB_USER, DB_NAME, DB_PASSWORD)
	URL := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", DB_USER, DB_PASSWORD, DB_HOST, DB_PORT, DB_NAME)
	// dsn := "root:root@tcp(127.0.0.1:3306)/btpn?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", URL)
	if err != nil {
		fmt.Printf("Cannot connect to %s database", DB_DRIVER)
		log.Fatal("This is the error:", err)
	} else {
		fmt.Printf("We are connected to the %s database", DB_DRIVER)
	}

	err = db.Debug().AutoMigrate(&models.User{}, &models.Photo{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	err = db.Debug().Model(&models.Photo{}).AddForeignKey("user_id", "users(id)", "cascade", "cascade").Error
	if err != nil {
		log.Fatalf("attaching foreign key error: %v", err)
	}

	return db
}
