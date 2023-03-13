package repository

import (
	"github.com/ervinismu/purplestore/models"
	"github.com/ervinismu/purplestore/schema"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return UserRepository{DB: db}
}

func (ur UserRepository) GetUsers() ([]models.User, error) {
	var users []models.User
	err := ur.DB.Find(&users).Error
	if err != nil {
		return users, err
	}

	return users, nil
}

func (ur UserRepository) ShowUser(id string) (models.User, error) {
	var user models.User
	if err := models.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserRepository) CreateUser(data schema.CreateUserReq) (models.User, error) {
	pass, _ := hashPassword(data.Password)
	user := models.User{
		Email: data.Email,
		Username: data.Username,
		Password: pass,
	}
	err := ur.DB.Create(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserRepository) UpdateUser(id string, data schema.UpdateUserReq) (models.User, error) {
	var user models.User
	if err := ur.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	user.Username = data.Username
	err := ur.DB.Save(&user).Error
	if err != nil {
		return user, err
	}

	return user, nil
}

func (ur UserRepository) DeleteUser(id string) (models.User, error) {
	var user models.User
	if err := ur.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}

	err := ur.DB.Delete(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 1)
	if err != nil {
		return "", err
	}

	return string(bytes), err
}
