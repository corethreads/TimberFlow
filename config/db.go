package config

import (
	"log"
	"os"
	"server/internal/auth/models/entity"

	"github.com/joho/godotenv" // 👈 new import
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// Load .env file into os.Environ
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ No .env file found, falling back to system env")
	}

	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		log.Fatal("❌ DATABASE_URL not set in .env or system env")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[ERROR] Failed to connect to database:", err)
	}

	if err := db.AutoMigrate(&entity.User{}); err != nil {
		log.Fatal("❌ Failed to migrate models:", err)
	}

	DB = db
	log.Println("✅ Database connected & migrated")
}
