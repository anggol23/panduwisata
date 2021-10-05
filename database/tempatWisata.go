package database

import (
	"mini-project/config"
	"mini-project/model"
)

func GetAllTempatWisata() []model.TempatWisata {
	var tempatWisata []model.TempatWisata
	config.DB.Find(&tempatWisata).Joins("Category")
	return tempatWisata
}

func GetTempatWisataByID(id string) model.TempatWisata {
	var tempatWisata model.TempatWisata
	config.DB.Where("id = ?", id).Find(&tempatWisata).Joins("Category")
	return tempatWisata
}

func CreateTempatWisata(tempatWisata model.TempatWisata) model.TempatWisata {
	config.DB.Create(&tempatWisata).Joins("Category")
	return tempatWisata
}
func DeleteTempatWisataByID(id string) {
	var tempatWisata model.TempatWisata
	config.DB.Where("id = ?", id).Delete(&tempatWisata)
}
func UpdateTempatWisataByID(id string, tempatWisata model.TempatWisata) {
	config.DB.Where("id = ?", id).Updates(&tempatWisata)
}
