package config

import (
	"log"
	"os"
	"sync"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	database *gorm.DB
	once     sync.Once
)

func LoadDB() {
	once.Do(func() {
		if err := godotenv.Load(); err != nil {
			log.Println("No .env file found")
		}

		dsn := os.Getenv("DATABASE_URL")
		if dsn == "" {
			log.Fatal("DATABASE_URL not found in environment")
		}

		var err error
		database, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal("Failed to connect to database:", err)
		}

		log.Println("[corethreads] Database connected successfully")
	})
}

// GetDB returns the database connection
func GetDB() *gorm.DB {
	if database == nil {
		log.Fatal("Database not initialized. Call LoadDB() first")
	}
	return database
}
