package model

import "gorm.io/gorm"

// import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"password,omitempty"`
	Phone_Numb string `json:"phone_numb"`
}
