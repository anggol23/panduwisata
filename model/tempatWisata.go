package model

import "gorm.io/gorm"

type TempatWisata struct {
	gorm.Model
	ID           uint        `gorm:"primarykey"`
	Name         string      `json:"name"`
	CategoriesID uint        `json:"category_id,omitempty"`
	Categories   []*Category `gorm:"many2many:categories" json:"categories,omitempty"`
	Description  string      `json:"description"`
	Location     string      `json:"location"`
	Ticket_Price string      `json:"ticket_price"`
}
