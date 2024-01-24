package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
	"github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/paymentintent"
)

func (api *API) Webhook(c *gin.Context) {
	var event stripe.Event

	if err := c.ShouldBindJSON(&event); err != nil {
		fmt.Println("Error binding event: ", err.Error())
		return
	}

	switch event.Type {
	case "payment_intent.succeeded":
		var paymentIntent stripe.PaymentIntent
		err := json.Unmarshal(event.Data.Raw, &paymentIntent)
		if err != nil {
			fmt.Println("Error unmarshalling payment intent: ", err.Error())
			return
		}

		productID, err := uuid.Parse(paymentIntent.Metadata["product_id"])
		if err != nil {
			fmt.Println("Error parsing product ID: ", err.Error())
			return
		}

		product, err := models.GetProductByID(api.db, productID)
		if err != nil {
			fmt.Println("Error parsing product ID: ", err.Error())
			return
		}

		user, err := models.CreateUser(api.db, &models.CreateUserRequest{
			FirstName: paymentIntent.Metadata["first_name"],
			LastName:  paymentIntent.Metadata["last_name"],
			Email:     paymentIntent.Metadata["email"],
			ZipCode:   paymentIntent.Metadata["zip_code"],
		})

		if err != nil {
			fmt.Println("Error creating user: ", err.Error())
			return
		}

		err = models.CreateTransaction(api.db, &models.CreateTransactionRequest{
			ProductID:                       product.ID,
			UserID:                          user.ID,
			Amount:                          product.Price,
			StripePaymentIntentID:           paymentIntent.ID,
			StripePaymentIntentClientSecret: paymentIntent.ClientSecret,
		})

		if err != nil {
			fmt.Println("Error creating transaction: ", err.Error())
			return
		}
	default:
		fmt.Println("Unhandled event type: ", event.Type)
		return
	}
}

type CreatePaymentIntentRequest struct {
	ProductID uuid.UUID `json:"product_id" binding:"required"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	ZipCode   string    `json:"zip_code"`
	IsJoining string    `json:"is_joining"`
}

func (api *API) CreatePaymentIntent(c *gin.Context) {
	var req CreatePaymentIntentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := models.GetProductByID(api.db, req.ProductID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paymentParams := &stripe.PaymentIntentParams{
		Amount:   stripe.Int64(int64(product.Price * 100)),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
		Metadata: map[string]string{
			"product_id": req.ProductID.String(),
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"email":      req.Email,
			"zip_code":   req.ZipCode,
			"is_joining": req.IsJoining,
		},
	}

	pi, err := paymentintent.New(paymentParams)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"client_secret": pi.ClientSecret})
}

type SubmitPaymentIntentRequest struct {
	PIsecret string `json:"pi_secret" binding:"required"`
}

func (api *API) SubmitPaymentIntent(c *gin.Context) {
	var req SubmitPaymentIntentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paymentParams := &stripe.PaymentIntentConfirmParams{
		ReturnURL: stripe.String("http://localhost:3000/"),
	}

	pi, err := paymentintent.Confirm(req.PIsecret, paymentParams)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"client_secret": pi.ClientSecret})
}

type UpdatePaymentIntentRequest struct {
	ClientSecret string    `json:"client_secret" binding:"required"`
	ProductID    uuid.UUID `json:"product_id" binding:"required"`
	FirstName    string    `json:"first_name" binding:"required"`
	LastName     string    `json:"last_name" binding:"required"`
	Email        string    `json:"email" binding:"required"`
	ZipCode      string    `json:"zip_code" binding:"required"`
	IsJoining    string    `json:"is_joining" binding:"required"`
}

func (api *API) UpdatePaymentIntent(c *gin.Context) {
	var req UpdatePaymentIntentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	product, err := models.GetProductByID(api.db, req.ProductID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	paymentParams := &stripe.PaymentIntentParams{
		Amount: stripe.Int64(int64(product.Price * 100)),
		Metadata: map[string]string{
			"product_id": req.ProductID.String(),
			"first_name": req.FirstName,
			"last_name":  req.LastName,
			"email":      req.Email,
			"zip_code":   req.ZipCode,
			"is_joining": req.IsJoining,
		},
	}

	pi, err := paymentintent.Update(req.ClientSecret, paymentParams)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"client_secret": pi.ClientSecret})
}
