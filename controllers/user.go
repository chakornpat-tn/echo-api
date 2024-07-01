package controllers

import (
	"echo-api/database"
	"echo-api/models"
	"fmt"
)

func CreateUser(userReq *models.User) error {
	result := database.DB.Create(userReq)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func FindUser(id uint32) (models.User, error) {
	user := models.User{}
	if err := database.DB.First(&user, id).Error; err != nil {
		return user, err
	}

	return user, nil
}

func ListUser(search string, page, perPage uint32) ([]models.User, error) {

	users := []models.User{}
	tw := database.DB.Table("users")
	if search != "" {
		tw = tw.Where("users ILIKE ?", "%"+search+"%", "OR users ILIKE ?", "%"+search+"%")
	}

	if err := tw.Limit(int(perPage)).Offset(int((page - 1) * perPage)).Scan(&users).Error; err != nil {
		return users, err
	}

	return users, nil
}

func UpdateUser(id uint32, userReq *models.User) error {

	user := models.User{}
	if err := database.DB.First(&user, id).Error; err != nil {
		return err
	}

	if err := database.DB.Model(&user).Updates(userReq).Error; err != nil {
		return err
	}

	return nil
}

func DeleteUser(id uint32) error {
	result := database.DB.Delete(&models.User{}, id)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("user with ID %d not found", id)
	}

	return nil
}
