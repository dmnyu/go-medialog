package database

import "github.com/dmnyu/go-medialog/models"

func FindUsers() (*[]models.User, error) {
	var users = []models.User{}
	if err := db.Find(&users).Error; err != nil {
		return nil, err
	}
	return &users, nil
}

func FindUser(id int) (*models.User, error) {
	var user = models.User{}
	if err := db.Where("id = ?", &user).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func CreateUser(user *models.User) error {
	if err := db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func DeleteUser(user *models.User) error {
	if err := db.Delete(user).Error; err != nil {
		return err
	}
	return nil
}

func UpdateUser(user *models.User) error {
	if err := db.Updates(user).Error; err != nil {
		return err
	}
	return nil
}
