package model

import "gorm.io/gorm"

type Transaction struct {
	gorm.Model
	ID             uint          `gorm:"primaryKey"`
	UserID         uint          `json:"user_id"`
	User           *User         `json:"user"`
	TempatWisataID uint          `json:"tempatwisata_id,"`
	TempatWisata   *TempatWisata `json:"tempatwisata,omitempty"`
}
