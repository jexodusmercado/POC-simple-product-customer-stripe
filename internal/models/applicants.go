package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Applicants struct {
	ID              uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	FirstName       string     `gorm:"type:varchar(255); not null;" json:"first_name"`
	LastName        string     `gorm:"type:varchar(255); not null;" json:"last_name"`
	Email           string     `gorm:"type:varchar(255); not null;" json:"email"`
	Phone           string     `gorm:"type:varchar(255); not null;" json:"phone"`
	ZipCode         string     `gorm:"type:varchar(255); not null;" json:"zip_code"`
	LinkedInProfile string     `gorm:"type:varchar(255); nullable;" json:"linked_in_profile"`
	JobTitle        string     `gorm:"type:varchar(255); not null;" json:"job_title"`
	ApplicationDate *time.Time `gorm:"nullable" json:"application_date,omitempty"`
	FileURL         string     `gorm:"type:varchar(255); nullable;" json:"file_url,omitempty"`
}

type ApplicantRequest struct {
	FirstName       string     `form:"first_name,omitempty" binding:"required"`
	LastName        string     `form:"last_name,omitempty" binding:"required"`
	Email           string     `form:"email,omitempty" binding:"required"`
	Phone           string     `form:"phone,omitempty" binding:"required"`
	ZipCode         string     `form:"zip_code,omitempty" binding:"required"`
	JobTitle        string     `form:"job_title,omitempty" binding:"required"`
	LinkedInProfile string     `form:"linked_in_profile,omitempty"`
	ApplicationDate *time.Time `form:"application_date,omitempty"`
	FileURL         string
}

func GetAllApplicants(tx *gorm.DB) ([]Applicants, error) {
	var applicants []Applicants
	if err := tx.Find(&applicants).Error; err != nil {
		return nil, err
	}
	return applicants, nil
}

func CreateApplicant(tx *gorm.DB, req ApplicantRequest) error {
	applicant := &Applicants{
		FirstName:       req.FirstName,
		LastName:        req.LastName,
		Email:           req.Email,
		Phone:           req.Phone,
		ZipCode:         req.ZipCode,
		LinkedInProfile: req.LinkedInProfile,
		JobTitle:        req.JobTitle,
		ApplicationDate: req.ApplicationDate,
        FileURL:         req.FileURL,
	}

	if err := tx.Create(applicant).Error; err != nil {
		return err
	}
	return nil
}

func GetApplicantById(tx *gorm.DB, id string) (*Applicants, error) {
	var applicant Applicants
	if err := tx.Where("id = ?", id).First(&applicant).Error; err != nil {
		return nil, err
	}
	return &applicant, nil
}
