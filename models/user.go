package models

import (
	"time"

	"gorm.io/gorm"
)



type User struct {
	Phone string  `gorm:"primaryKey"`
	Password string
	Balance string
	Token string
	Name string
	ProfileID string
	Bonus string 
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}