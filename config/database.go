package config

import (
	"fmt"
	"log"
	"os"
	"time"

	"backend_article/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"github.com/joho/godotenv"
)

var DB *gorm.DB

func ConnectDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var database *gorm.DB
	for i := 0; i < 3; i++ {
		database, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Println("Gagal menghubungkan ke database:", err)
		log.Println("Mencoba kembali dalam 3 detik")
		time.Sleep(3 * time.Second)
	}

	if err != nil {
		log.Fatal("Gagal menghubungkan ke database setelah beberapa kali percobaan:", err)
	}

	err = database.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Gagal melakukan migrasi database:", err)
	}

	DB = database
	log.Println("Berhasil terhubung ke database!")
}
