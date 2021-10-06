package database

import (
	"errors"
	"mini-project/config"
	"mini-project/model"
)

// func IsValidBasic(u string) bool {
// 	return true
// }

// func IsValid(u string, p string) bool {
// 	var user model.User
// 	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
// 		return false
// 	}

// 	return p == user.Password
// }
// func GetAllUser() []model.User {
// 	var user []model.User
// 	config.DB.Find(&user)
// 	return user
// }

// func GetUserByID(id string) model.User {
// 	var user model.User
// 	config.DB.Where("id = ?", id).Find(&user)
// 	return user
// }

// func CreateUser(user model.User) model.User {
// 	config.DB.Create(&user)
// 	return user
// }
// func DeleteUserByID(id string) {
// 	var user model.User
// 	config.DB.Where("id = ?", id).Delete(&user)
// }
// func UpdateUserByID(id string, user model.User) {
// 	config.DB.Where("id = ?", id).Updates(&user)
// }

func GetAllUser() []model.User {
	var user []model.User
	config.DB.Find(&user)
	return user
}

func GetUserByID(id string) model.User {
	var user model.User
	config.DB.Where("id = ?", id).Find(&user)
	return user
}

func CreateUser(user model.User) model.User {
	config.DB.Create(&user)
	return user
}

func UpdateUser(id string, user model.User) model.User {
	config.DB.Where("id = ?", id).Updates(&user)
	return user
}

func IsValid(u string, p string) (model.User, error) {
	var user model.User
	if err := config.DB.Where("email = ?", u).Find(&user).Error; err != nil {
		// fmt.Println(err)
		return model.User{}, err

	}
	if user.Password != p {
		return model.User{}, errors.New("email or password incorrect")
	}

	return user, nil
}
