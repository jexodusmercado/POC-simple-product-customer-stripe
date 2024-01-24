package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	FirstName string `gorm:"type:varchar(255)" json:"first_name,omitempty"`
	LastName string `gorm:"type:varchar(255)" json:"last_name,omitempty"`
	Email string `gorm:"type:varchar(255)" json:"email,omitempty"`
	ZipCode string `gorm:"type:varchar(255)" json:"zip_code,omitempty"`
}

type CreateUserRequest struct {
	FirstName string `json:"first_name,omitempty" binding:"required"`
	LastName string `json:"last_name,omitempty" binding:"required"`
	Email string `json:"email,omitempty" binding:"required"`
	ZipCode string `json:"zip_code,omitempty" binding:"required"`
}

func CreateUser(tx *gorm.DB, req *CreateUserRequest) (User, error) {

	user := User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		ZipCode: req.ZipCode,
	}

	err := tx.Create(&user).Error

	return user, err
}

func GetUsers(tx *gorm.DB) ([]User, error) {

	var users []User

	err := tx.Find(&users).Error

	return users, err

}

func GetUserByID(tx *gorm.DB, id uuid.UUID) (User, error) {

	var user User

	err := tx.Where("id = ?", id).First(&user).Error

	return user, err

}

func UpdateUser(tx *gorm.DB, id uuid.UUID, req *CreateUserRequest) error {

	user := User{
		FirstName: req.FirstName,
		LastName: req.LastName,
		Email: req.Email,
		ZipCode: req.ZipCode,
	}

	return tx.Model(&user).Where("id = ?", id).Updates(&user).Error

}