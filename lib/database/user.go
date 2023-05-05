package database

import (
	"praktikum/config"
	"praktikum/models"
)

// Get all user
func GetUser() (users []models.User, err error) {
	err = config.DB.Preload("Blog").Find(&users).Error

	if err != nil {
		return []models.User{}, err
	}

	//make to return blog data

	return
}

// Create user
func CreateUser(user models.User) (models.User, error) {
	err := config.DB.Create(&user).Error

	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

// Get user by id
func GetUserById(id any) (models.User, error) {
	var user models.User

	err := config.DB.Preload("Blog").Where("id = ?", id).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Update user by id
func UpdateUser(user models.User, id any) (models.User, error) {
	err := config.DB.Where("id = ?", id).Updates(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Delete user by id
func DeleteUser(id any) (interface{}, error) {
	err := config.DB.Where("id = ?", id).Delete(&models.User{}).Error

	if err != nil {
		return nil, err
	}

	return "success delete user", nil
}

//JWT
func LoginUser(user models.User) (models.User, error) {

	err := config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(&user).Error

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}