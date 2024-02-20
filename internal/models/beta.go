package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Beta struct {
	ID          	uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id,omitempty"`
	FirstName 		string `gorm:""type:varchar(255);" json:"first_name,omitempty"`
	LastName 		string `gorm:""type:varchar(255);" json:"last_name,omitempty"`
	Email  			string `gorm:""type:varchar(255);" json:"email,omitempty"`
	Phone			string `gorm:""type:varchar(255);" json:"phone,omitempty"`
	Zipcode       	string `gorm:""type:varchar(255);" json:"zip_code,omitempty"`
}

type CreateBetaUserRequest struct {
	FirstName       string `json:"first_name,omitempty" binding:"required"`
	LastName    	string `json:"last_name,omitempty" binding:"required"`
	Email       	string `json:"email,omitempty" binding:"required"`
	Phone       	string `json:"phone,omitempty" binding:"required"`
	Zipcode       	string `json:"zip_code,omitempty" binding:"required"`
}

func CreateBetaUser(tx *gorm.DB, req CreateBetaUserRequest) (error) {

	betaUser := &Beta{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Zipcode:   req.Zipcode,
	}

	if err := tx.Create(betaUser).Error; err != nil {
		return err
	}
	return nil
}

func GetBetaUsers(tx *gorm.DB) ([]Beta, error) {

	var betaUsers []Beta

	err := tx.Find(&betaUsers).Error

	return betaUsers, err

}

func GetBetaUserById(tx *gorm.DB, id string) (*Beta, error) {
	var betaUser Beta
	if err := tx.Where("id = ?", id).First(&betaUser).Error; err != nil {
		return nil, err
	}
	return &betaUser, nil

}

func GetBetaUserByEmail(tx *gorm.DB, email string) (*Beta, error) {
	var betaUser Beta
	if err := tx.Where("email = ?", email).First(&betaUser).Error; err != nil {
		return nil, err
	}
	return &betaUser, nil

}
