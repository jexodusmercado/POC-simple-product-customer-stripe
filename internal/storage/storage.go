package storage

import (
	"fmt"

	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Dial(env conf.GlobalConfig) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s",
		env.DB_HOST,
		env.DB_USER,
		env.DB_PASSWORD,
		env.DB_NAME,
		env.DB_PORT,
	)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		TranslateError:                           true,
	})

	if err != nil {
		panic("Failed to connect to database")
	}

	return database
}

func MigrateDatabase(tx *gorm.DB) {
	var err error
	err = tx.AutoMigrate(&models.Product{})
	if err != nil {
		panic("Failed to migrate product")
	}

	err = tx.AutoMigrate(&models.Transaction{})
	if err != nil {
		panic("Failed to migrate customer")
	}

	err = tx.AutoMigrate(&models.User{})
	if err != nil {
		panic("Failed to migrate user")
	}

	err = tx.AutoMigrate(&models.Inquiry{})
	if err != nil {
		panic("Failed to migrate inquiry")
	}

}
