package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	Name            string    `gorm:"type:varchar(255)" json:"name,omitempty"`
	Description     string    `gorm:"type:varchar(255)" json:"description,omitempty"`
	BasePrice       float64   `gorm:"type:float" json:"basePrice,omitempty"`
	DiscountedPrice float64   `gorm:"type:float" json:"discountedPrice,omitempty"`
	Quantity        int       `gorm:"type:int" json:"quantity,omitempty"`
}

type CreateProductRequest struct {
	Name            string  `json:"name,omitempty" binding:"required"`
	Description     string  `json:"description,omitempty"`
	BasePrice       float64 `json:"basePrice,omitempty" binding:"required"`
	DiscountedPrice float64 `json:"discountedPrice,omitempty"`
	Quantity        int     `json:"quantity,omitempty" binding:"required"`
}

func CreateProduct(tx *gorm.DB, req *CreateProductRequest) error {

	product := Product{
		Name:            req.Name,
		Description:     req.Description,
		BasePrice:       req.BasePrice,
		DiscountedPrice: req.DiscountedPrice,
		Quantity:        req.Quantity,
	}

	return tx.Create(&product).Error

}

func GetProducts(tx *gorm.DB) ([]Product, error) {

	var products []Product

	if err := tx.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil

}

func GetProductByID(tx *gorm.DB, id uuid.UUID) (Product, error) {

	var product Product

	err := tx.Where("id = ?", id).First(&product).Error

	return product, err

}

func UpdateProduct(tx *gorm.DB, id uuid.UUID, req *CreateProductRequest) error {

	product := Product{
		Name:            req.Name,
		Description:     req.Description,
		BasePrice:       req.BasePrice,
		DiscountedPrice: req.DiscountedPrice,
		Quantity:        req.Quantity,
	}

	return tx.Model(&product).Where("id = ?", id).Updates(product).Error
}
