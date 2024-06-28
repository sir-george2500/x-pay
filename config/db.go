package config

import (
    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "log"
    "os"
)

var DB *gorm.DB

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Construct DSN using environment variables
    dsn := "postgresql://" + os.Getenv("DB_USER") + ":" + os.Getenv("DB_PASSWORD") +
           "@" + os.Getenv("DB_HOST") + ":" + os.Getenv("DB_PORT") + "/" +
           os.Getenv("DB_NAME")

    var errDB error
    DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if errDB != nil {
        log.Fatalf("Failed to connect to database: %v", errDB)
    }
}

func PingDB() error {
    sqlDB, err := DB.DB()
    if err != nil {
        return err
    }
    return sqlDB.Ping()
}

