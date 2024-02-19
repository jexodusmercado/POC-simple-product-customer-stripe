package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
