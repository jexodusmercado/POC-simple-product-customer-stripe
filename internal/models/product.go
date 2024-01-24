package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Product struct {
	ID uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"ID,omitempty"`
	Name string `gorm:"type:varchar(255)" json:"name,omitempty"`
	Description string `gorm:"type:varchar(255)" json:"description,omitempty"`
	Price float64 `gorm:"type:float" json:"price,omitempty"`
	Quantity int `gorm:"type:int" json:"quantity,omitempty"`
}

type CreateProductRequest struct {
	Name string `json:"name,omitempty" binding:"required"`
	Description string `json:"description,omitempty"`
	Price float64 `json:"price,omitempty" binding:"required"`
	Quantity int `json:"quantity,omitempty" binding:"required"`
}

func CreateProduct(tx *gorm.DB, req *CreateProductRequest) error {

	product := Product{
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Quantity: req.Quantity,
	}

	return tx.Create(&product).Error

}

func GetProducts(tx *gorm.DB) ([]Product, error) {

	var products []Product

	err := tx.Find(&products).Error

	return products, err

}

func GetProductByID(tx *gorm.DB, id uuid.UUID) (Product, error) {

	var product Product

	err := tx.Where("id = ?", id).First(&product).Error

	return product, err

}

func UpdateProduct(tx *gorm.DB, id uuid.UUID, req *CreateProductRequest) error {

	product := Product{
		Name: req.Name,
		Description: req.Description,
		Price: req.Price,
		Quantity: req.Quantity,
	}

	return tx.Model(&product).Where("id = ?", id).Updates(product).Error

}