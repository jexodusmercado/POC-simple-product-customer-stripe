package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Transaction struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	ProductID uuid.UUID `gorm:"type:uuid;" json:"product_id,omitempty"`
	UserID uuid.UUID `gorm:"type:uuid;" json:"user_id,omitempty"`
	Amount float64 `gorm:"type:float;" json:"amount,omitempty"`
	StripePaymentIntentID string `json:"stripe_payment_intent_id,omitempty"`
	StripePaymentIntentClientSecret string `json:"stripe_payment_intent_client_secret,omitempty"`
}

type CreateTransactionRequest struct {
	ProductID uuid.UUID `json:"product_id,omitempty" binding:"required"`
	UserID uuid.UUID `json:"user_id,omitempty" binding:"required"`
	Amount float64 `json:"amount,omitempty" binding:"required"`
	StripePaymentIntentID string `json:"stripe_payment_intent_id,omitempty" binding:"required"`
	StripePaymentIntentClientSecret string `json:"stripe_payment_intent_client_secret,omitempty" binding:"required"`
}

func CreateTransaction(tx *gorm.DB, req *CreateTransactionRequest) (Transaction, error) {
	
	transaction := Transaction{
		ProductID: req.ProductID,
		UserID: req.UserID,
		Amount: req.Amount,
		StripePaymentIntentID: req.StripePaymentIntentID,
		StripePaymentIntentClientSecret: req.StripePaymentIntentClientSecret,
	}

	if err := tx.Create(&transaction).Error; err != nil {
		return Transaction{}, err
	}

	return transaction, nil

}

func GetTransactions(tx *gorm.DB) ([]Transaction, error) {
	
	var transactions []Transaction

	err := tx.Find(&transactions).Error

	return transactions, err

}

func GetTransactionByID(tx *gorm.DB, id string) (Transaction, error) {
	
	var transaction Transaction

	err := tx.Where("id = ?", id).First(&transaction).Error

	return transaction, err

}

func GetTransactionByStripePaymentIntentID(tx *gorm.DB, id string) (Transaction, error) {
	
	var transaction Transaction

	err := tx.Where("stripe_payment_intent_id = ?", id).First(&transaction).Error

	return transaction, err

}

func UpdateTransaction(tx *gorm.DB, id string, req *CreateTransactionRequest) error {
	
	transaction := Transaction{
		ProductID: req.ProductID,
		UserID: req.UserID,
		Amount: req.Amount,
		StripePaymentIntentID: req.StripePaymentIntentID,
		StripePaymentIntentClientSecret: req.StripePaymentIntentClientSecret,
	}

	return tx.Save(&transaction).Error

}
