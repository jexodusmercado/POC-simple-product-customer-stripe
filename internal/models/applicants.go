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
}

type ApplicantRequest struct {
	FirstName           string     `json:"first_name,omitempty" binding:"required"`
	LastName            string     `json:"last_name,omitempty" binding:"required"`
	Email               string     `json:"email,omitempty" binding:"required"`
	Phone               string     `json:"phone,omitempty" binding:"required"`
	ZipCode             string     `json:"zip_code,omitempty" binding:"required"`
	JobTitle            string     `json:"job_title,omitempty" binding:"required"`
	LinkedInProfile     string     `json:"linked_in_profile,omitempty"`
	ApplicationDate     *time.Time `json:"application_date,omitempty"`
	ApplicantAttachment []byte     `json:"applicant_attachment,omitempty" binding:"required"`
	ApplicantFileName   string     `json:"applicant_filename,omitempty" binding:"required"`
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
