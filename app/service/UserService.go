package service

import (
	"app/models"
	"app/repository"
)

func LoginUser(email string, passwordHash string) (*models.User, error) {
	user, err := repository.LoginUser(email, passwordHash)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserById(id int) (*models.User, error) {
	user, err := repository.GetUserById(id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user, err := repository.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func Adduser(newUser models.User) (*models.User, error) {
	user, err := repository.Adduser(newUser)
	if err == nil {
		return nil, err
	}

	return user, nil
}
