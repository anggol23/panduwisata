package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID             uint            `gorm:"primarykey"`
	UserID         uint            `json:"user_id"`
	User           []*User         `json:"User"`
	TempatWisataID uint            `json:"tempatWisata_id"`
	TempatWisata   []*TempatWisata `json:"tempatWisata"`
}
