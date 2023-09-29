package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID int `json:"id"`
	Name    string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Password    string `json:"password"`
	gorm.Model
}