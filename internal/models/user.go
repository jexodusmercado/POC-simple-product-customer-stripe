package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	FirstName   string     `gorm:"type:varchar(255); not null;" json:"first_name,omitempty"`
	LastName    string     `gorm:"type:varchar(255); not null;" json:"last_name,omitempty"`
	Email       string     `gorm:"type:varchar(255); unique; not null;" json:"email,omitempty"`
	ZipCode     string     `gorm:"type:varchar(255); not null;" json:"zip_code,omitempty"`
	PhoneNumber string     `gorm:"type:varchar(255)" json:"phone_number,omitempty"`
	IsJoinBeta  *time.Time `gorm:"nullable" json:"is_join_beta,omitempty"`
}

type CreateUserRequest struct {
	FirstName   string     `json:"first_name,omitempty" binding:"required"`
	LastName    string     `json:"last_name,omitempty" binding:"required"`
	Email       string     `json:"email,omitempty" binding:"required"`
	ZipCode     string     `json:"zip_code,omitempty" binding:"required"`
	PhoneNumber string     `json:"phone_number,omitempty"`
	IsJoinBeta  *time.Time `json:"is_join_beta,omitempty"`
}

func CreateUser(tx *gorm.DB, req *CreateUserRequest) (User, error) {

	user := User{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		ZipCode:     req.ZipCode,
		IsJoinBeta:  req.IsJoinBeta,
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

func GetUserByEmail(tx *gorm.DB, email string) (User, error) {
	var user User

	err := tx.Where("email = ?", email).First(&user).Error

	return user, err

}

func UpdateUser(tx *gorm.DB, id uuid.UUID, req *CreateUserRequest) error {

	user := User{
		FirstName:  req.FirstName,
		LastName:   req.LastName,
		Email:      req.Email,
		ZipCode:    req.ZipCode,
		IsJoinBeta: req.IsJoinBeta,
	}

	return tx.Model(&user).Where("id = ?", id).Updates(&user).Error

}

func CheckUserExists(tx *gorm.DB, email string) (bool, error) {
	var user User
	err := tx.Where("email = ?", email).First(&user).Error

	if err != nil {
		return false, err
	}

	return true, nil
}
