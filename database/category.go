package database

import (
	"mini-project/config"
	"mini-project/model"
)

// func GetAllCategory() []model.Category {
// 	var category []model.Category
// 	config.DB.Joins("TempatWisata").Find(&category)
// 	return category
// }

// func GetCategoryByID(id string) model.Category {
// 	var category model.Category
// 	config.DB.Where("id = ?", id).Joins("TempatWisata").Find(&category)
// 	return category
// }

// func CreateCategory(category model.Category) model.Category {
// 	config.DB.Create(&category).Joins("TempatWisata")
// 	return category
// }
// func DeleteCategoryByID(id string) {
// 	var category model.Category
// 	config.DB.Where("id = ?", id).Delete(&category)
// }
// func UpdateCategoryByID(id string, Category model.Category) {
// 	config.DB.Where("id = ?", id).Updates(&Category).Joins("TempatWisata")
// }

func GetAllCategory() []model.Category {
	var category []model.Category
	config.DB.Find(&category)
	return category
}

func GetCategoryByID(id string) model.Category {
	var category model.Category
	config.DB.Where("id = ?", id).Find(&category)
	return category
}

func CreateCategory(category model.Category) model.Category {
	config.DB.Create(&category)
	return category
}
func DeleteCategoryByID(id string) {
	var category model.Category
	config.DB.Where("id = ?", id).Delete(&category)
}
func UpdateCategory(id string, Category model.Category) {
	config.DB.Where("id = ?", id).Updates(&Category)
}
