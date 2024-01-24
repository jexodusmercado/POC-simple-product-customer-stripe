package main

import (
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/api"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/storage"
	"github.com/stripe/stripe-go/v76"

	"github.com/gin-gonic/gin"
)

func main() {
	config := conf.InitEnv()

	stripe.Key = config.STRIPE_SECRET_KEY

	if config.APPLICATION_ENV == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	router := gin.Default()
	db := storage.Dial(config)
	storage.MigrateDatabase(db)
	server := api.NewAPIWithVersion(router, db, &config)
	server.Run()
}
