package repository

import (
	"app/helper"
	"app/models"
	"errors"

	"gorm.io/gorm"
)

// type UserRepository interface {
// 	loginUser(email string, passwordHash string) (*models.User, error)
// 	getUserById(id int) (*models.User, error)
// 	adduser(user models.User) (*models.User, error)
// }

func LoginUser(email string, passwordHash string) (*models.User, error) {
	var user models.User

	result := helper.DBConn.Where("email = ? AND password_hash = ?", email, passwordHash).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("Login user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func GetUserById(id int) (*models.User, error) {
	var user models.User

	result := helper.DBConn.First(&user, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User

	result := helper.DBConn.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &user, nil
}

func Adduser(user models.User) (*models.User, error) {
	resultCreate := helper.DBConn.Create(&user)
	if resultCreate.Error == nil {
		return nil, resultCreate.Error
	}

	return &user, nil
}
