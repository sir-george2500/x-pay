package models

import (
    "time"
)

type User struct {
    ID           uint      `gorm:"primaryKey"`
    Username     string    `gorm:"uniqueIndex;not null"`
    FullName     string    `gorm:"not null"`
    Email        string    `gorm:"uniqueIndex;not null"`
    Password     string    `gorm:"not null"`
    Address      string    `gorm:"not null"`
    DateOfBirth  time.Time `gorm:"not null"`
    PhoneNumber  string    `gorm:"uniqueIndex;not null"`
    AccountType  string    `gorm:"not null"`
    CreatedAt    time.Time `gorm:"autoCreateTime"`
    UpdatedAt    time.Time `gorm:"autoUpdateTime"`
}

