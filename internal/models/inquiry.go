package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Inquiry struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName string    `gorm:"type:varchar(255); not null;" json:"first_name"`
	LastName  string    `gorm:"type:varchar(255); not null;" json:"last_name"`
	Email     string    `gorm:"type:varchar(255); not null;" json:"email"`
	Phone     string    `gorm:"type:varchar(255); not null;" json:"phone"`
	Message   string    `gorm:"type:text; not null;" json:"message"`
}

type InquiryRequest struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Phone     string `json:"phone" binding:"required"`
	Message   string `json:"message" binding:"required"`
}

func GetAllInquiries(tx *gorm.DB) ([]Inquiry, error) {
	var inquiries []Inquiry
	if err := tx.Find(&inquiries).Error; err != nil {
		return nil, err
	}
	return inquiries, nil
}

func CreateInquiry(tx *gorm.DB, req InquiryRequest) error {
	inquiry := &Inquiry{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Message:   req.Message,
	}

	if err := tx.Create(inquiry).Error; err != nil {
		return err
	}
	return nil
}

func GetInquiryByID(tx *gorm.DB, id string) (*Inquiry, error) {
	var inquiry Inquiry
	if err := tx.Where("id = ?", id).First(&inquiry).Error; err != nil {
		return nil, err
	}
	return &inquiry, nil
}
