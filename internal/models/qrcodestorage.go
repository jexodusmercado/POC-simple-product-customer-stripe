package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type QrCodes struct {
	ID          	uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	TransactionID 	uuid.UUID `gorm:"type:uuid;" json:"transaction_id,omitempty"`
	UserID 			uuid.UUID `gorm:"type:uuid;" json:"user_id,omitempty"`
	S3Url         	string    `gorm:"type:longtext;not null;" json:"s3_url,omitempty"`
}

type CreateQrCodeRequest struct {
	TransactionID   uuid.UUID `json:"transaction_id,omitempty" binding:"required"`
	UserID    		uuid.UUID `json:"user_id,omitempty" binding:"required"`
	S3Url       	string `json:"s3_url,omitempty" binding:"required"`
}

func CreateQrCode(tx *gorm.DB, req *CreateQrCodeRequest) (QrCodes, error) {

	qrcode := QrCodes{
		TransactionID: 	req.TransactionID,
		UserID:  		req.UserID,
		S3Url:     		req.S3Url,
	}

	err := tx.Create(&qrcode).Error

	return qrcode, err
}

func GetQrCodes(tx *gorm.DB) ([]QrCodes, error) {

	var qrcodes []QrCodes

	err := tx.Find(&qrcodes).Error

	return qrcodes, err

}

func GetQrCodeByUserIdAndTransactionId(tx *gorm.DB, userID, transactionID uuid.UUID) (QrCodes, error) {

	var qrcode QrCodes
    err := tx.Where("user_id = ? AND transaction_id = ?", userID, transactionID).First(&qrcode).Error
    return qrcode, err

}

