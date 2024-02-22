package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QrCodes struct {
	ID            uuid.UUID  `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	TransactionID uuid.UUID  `gorm:"type:uuid;" json:"transaction_id,omitempty"`
	UserID        uuid.UUID  `gorm:"type:uuid;" json:"user_id,omitempty"`
	S3Url         string     `gorm:"type:text;not null;" json:"s3_url,omitempty"`
	IsQrUsed      *time.Time `gorm:"nullable" json:"is_qr_used,omitempty"`
}

type CreateQrCodeRequest struct {
	TransactionID uuid.UUID `json:"transaction_id,omitempty" binding:"required"`
	UserID        uuid.UUID `json:"user_id,omitempty" binding:"required"`
	S3Url         string    `json:"s3_url,omitempty" binding:"required"`
}

func CreateQrCode(tx *gorm.DB, req *CreateQrCodeRequest) (QrCodes, error) {

	qrcode := QrCodes{
		TransactionID: req.TransactionID,
		UserID:        req.UserID,
		S3Url:         req.S3Url,
	}

	err := tx.Create(&qrcode).Error
	if err != nil {
		return QrCodes{}, err
	}

	return qrcode, nil
}

func GetQrCodes(tx *gorm.DB) ([]QrCodes, error) {

	var qrcodes []QrCodes

	err := tx.Find(&qrcodes).Error
	if err != nil {
		return nil, err
	}

	return qrcodes, nil

}

func GetQrCodeByUserIDAndTransactionID(tx *gorm.DB, transactionID uuid.UUID, userID uuid.UUID) (QrCodes, error) {

	var qrcode QrCodes
	err := tx.Where("transaction_id = ? AND user_id = ?", transactionID, userID).First(&qrcode).Error
	if err != nil {
		return qrcode, err
	}

	return qrcode, nil

}
