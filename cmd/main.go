package main

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/api"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/storage"
	"github.com/stripe/stripe-go/v76"

	"github.com/gin-gonic/gin"
)

func main() {
	c := conf.InitEnv()

	stripe.Key = c.STRIPE_SECRET_KEY

	if c.APPLICATION_ENV == "development" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic("unable to load SDK config, " + err.Error())
	}

	// Create an Amazon S3 service client
	s3Client := s3.NewFromConfig(cfg)

	router := gin.Default()
	db := storage.Dial(c)
	storage.MigrateDatabase(db)
	server := api.NewAPIWithVersion(router, db, &c, s3Client)
	server.Run()
}
