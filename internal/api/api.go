package api

import (
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type API struct {
	handler  *gin.Engine
	db       *gorm.DB
	config   *conf.GlobalConfig
	s3Client *s3.Client
}

func NewAPI(handler *gin.Engine, db *gorm.DB, config *conf.GlobalConfig, s3Client *s3.Client) *API {
	return NewAPIWithVersion(handler, db, config, s3Client)
}

func NewAPIWithVersion(handler *gin.Engine, db *gorm.DB, conf *conf.GlobalConfig, s3Client *s3.Client) *API {
	api := &API{
		handler:  handler,
		db:       db,
		config:   conf,
		s3Client: s3Client,
	}

	//cors config
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Content-Type", "Authorization") // Add Authorization header
	corsConfig.AddAllowMethods("GET", "POST", "PATCH")

	api.handler.Use(cors.New(corsConfig))
	api.handler.Use(middleware.AddRequestID(api.config))

	test := api.handler.Group("test")

	test.GET("/email", api.testEmail)
	//test.GET("/qrcode", api.testQRCode)

	email := api.handler.Group("email")

	email.POST("/applicant", api.SendApplicantEmail)
	email.POST("/beta", api.SendBetaRegistrationEmail)
	email.POST("/contact", api.SendContactUsEmail)
	//email.POST("/qr", api.SendQrCodeEmail)

	payment := api.handler.Group("payments")

	payment.POST("/create-payment-intent", api.CreatePaymentIntent)
	payment.POST("/submit-payment-intent", api.SubmitPaymentIntent)
	payment.POST("/update-payment-intent", api.UpdatePaymentIntent)
	payment.POST("/webhook", api.Webhook)

	user := api.handler.Group("users")

	user.POST("", api.CreateUser)
	user.GET("", api.GetUsers)
	user.GET("/:id", api.GetUser)
	user.GET("/exist/:email", api.CheckUserExists)

	product := api.handler.Group("products")

	product.POST("", api.CreateProduct)
	product.GET("", api.GetProducts)
	product.GET("/:id", api.GetProduct)

	transaction := api.handler.Group("transactions")

	transaction.GET("", api.GetTransactions)
	transaction.GET("/:id", api.GetTransaction)

	inquiry := api.handler.Group("inquiries")

	inquiry.POST("", api.CreateInquiry)
	inquiry.GET("", api.GetInquiries)

	return api
}

func (api *API) Run() {
	api.handler.Run(":8080")
}
