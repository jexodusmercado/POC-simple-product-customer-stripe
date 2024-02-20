package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *API) SendQr(c *gin.Context) {
	mail := Mail{
		Subject:       "Test Email",
		CustomerEmail: "blueandraedevera@gmail.com",
		CustomerName:  "Blue Andrae",
		Body:          "<h1>This is a test email</h1>",
	}

	api.SendMail(mail)

	c.JSON(http.StatusOK, gin.H{
		"message": "Email sent",
	})
}

func (api *API) SendApplicantEmail(c *gin.Context) {
	var req Applicant
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := api.SendApplicationMail(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Email sent",
    })
}

func (api *API) SendBetaRegistrationEmail(c *gin.Context) {
	var req BetaList
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := api.SendBetaMail(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Email sent",
    })
}

func (api *API) SendContactUsEmail(c *gin.Context) {
	var req ContactUs
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := api.SendContactUsMail(req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Email sent",
    })
}


func (api *API) SendQrCodeEmail(c *gin.Context) {
	var req QrCode
    if err := c.ShouldBindJSON(&req); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := api.SendQrCodeMail(api.db, req); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "message": "Email sent",
    })
}