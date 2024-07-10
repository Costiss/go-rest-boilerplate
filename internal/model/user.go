package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    string `gorm:"primaryKey;type:uuid;default:gen_random_uuid()"`
	Name  string `gorm:"not null"`
	Email string `gorm:"unique;not null"`
}
