package model

import "gorm.io/gorm"

type TempatWisata struct {
	gorm.Model
	ID           uint      `gorm:"primarykey"`
	Name         string    `json:"name"`
	CategoryID   uint      `json:"category_id,omitempty"`
	Category     *Category `json:"category,omitempty"`
	UserID       uint      `json:"user_id"`
	User         *User     `json:"user,omitempty"`
	Description  string    `json:"description"`
	Location     string    `json:"location"`
	Ticket_Price string    `json:"ticket_price"`
}
