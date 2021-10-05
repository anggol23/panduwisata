package model

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	ID             uint            `gorm:"primarykey"`
	Name           string          `json:"name"`
	TempatWisataID uint            `json:"tempatwisata_id"`
	TempatWisata   []*TempatWisata `gorm:"many2many:tempatWisata"`
}
