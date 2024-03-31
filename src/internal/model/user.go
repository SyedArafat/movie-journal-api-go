package model

import (
	"time"
)

type User struct {
	ID              uint64     `gorm:"primaryKey;autoIncrement"`
	Name            string     `gorm:"not null"`
	Email           string     `gorm:"not null;unique"`
	EmailVerifiedAt *time.Time `gorm:"null"`
	Password        string     `gorm:"not null"`
	DateOfBirth     *time.Time `gorm:"null"`
	RememberToken   *string    `gorm:"null"`
	CreatedAt       time.Time  `gorm:"null"`
	UpdatedAt       time.Time  `gorm:"null"`
}
