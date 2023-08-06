package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() error {
	err := godotenv.Load("./../../.env")
	if err != nil {
		return fmt.Errorf("failed to load .env file: %w", err)
	}
	return nil
}

func InitDB(service string) error {

	err := LoadEnv()
	if err != nil {
		return err
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv(fmt.Sprintf("DB_USER_%s", service)),
		os.Getenv(fmt.Sprintf("DB_PASS_%s", service)),
		os.Getenv(fmt.Sprintf("DB_HOST_%s", service)),
		os.Getenv(fmt.Sprintf("DB_PORT_%s", service)),
		os.Getenv(fmt.Sprintf("DB_NAME_%s", service)))

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}

	DB = db
	return nil
}

func ConnectDB(service string) {
	if err := InitDB(service); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	fmt.Printf("Connected to the %s database!", service)
}
