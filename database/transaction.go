package database

import (
	"mini-project/config"
	"mini-project/model"
)

// import (
// 	"mini-project/config"
// 	"mini-project/model"
// )

// func GetAllTransaction() []model.Transaction {
// 	var transaction []model.Transaction
// 	config.DB.Find(&transaction).Joins("User", "TempatWisata")
// 	return transaction
// }

// func GetTransactionByID(id string) model.Transaction {
// 	var transaction model.Transaction
// 	config.DB.Where("id = ?", id).Find(&transaction).Joins("User", "TempatWisata")
// 	return transaction
// }

// func CreateTransaction(transaction model.Transaction) model.Transaction {
// 	config.DB.Create(&transaction)
// 	config.DB.Where("transactions.id = ?", transaction.ID).Joins("User").Joins("TempatWisata").Find(&transaction)
// 	return transaction
// }
// func DeleteTransactionByID(id string) {
// 	var transaction model.Transaction
// 	config.DB.Where("id = ?", id).Delete(&transaction)
// }
// func UpdateTransactionByID(id string, transaction model.Transaction) {
// 	config.DB.Where("id = ?", id).Updates(&transaction).Joins("USer").Joins("TempatWisata")
// }

func GetAllTransaction() []model.Transaction {
	var transaction []model.Transaction
	config.DB.Find(&transaction)
	return transaction
}

func GetTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Preload("User").Preload("TempatWisata.Category").Find(&transaction)
	return transaction
}

func CreateTransaction(transaction model.Transaction) model.Transaction {
	config.DB.Create(&transaction)
	config.DB.Where("transactions.id = ?", transaction.ID).Find(&transaction)
	return transaction
}

func UpdateTransaction(id string, transaction model.Transaction) model.Transaction {
	config.DB.Where("id = ?", id).Updates(&transaction).Joins("TempatWisata")
	return transaction
}
func DeleteTransactionByID(id string) model.Transaction {
	var transaction model.Transaction
	config.DB.Where("id = ?", id).Delete(&transaction)
	return transaction
}
