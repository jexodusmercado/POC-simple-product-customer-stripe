package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/helper"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) testEmail(c *gin.Context) {
	mail := Mail{
		Subject:       "Test Email",
		CustomerEmail: "jexodusmercado@gmail.com",
		CustomerName:  "Exo Mercado",
		Body:          "<h1>This is a test email</h1>",
	}

	api.SendMail(mail)

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent",
	})
}

func (api *API) testQRCode(c *gin.Context) {
	email := c.Query("email")

	user, err := models.GetUserByEmail(api.db, email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error getting user",
		})
		return
	}

	qrCodeDetails := helper.QRCodeDetails{
		UserID: user.ID.String(),
	}

	qrCode, err := helper.GenerateQRCode(qrCodeDetails)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error generating QR code",
		})
		return
	}

	key, err := api.UploadQRCode(qrCode, user.ID.String())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error uploading QR code",
            "error": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "QR code uploaded",
		"key":     key,
	})
}
