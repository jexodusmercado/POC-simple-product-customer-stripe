package api

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateInquiry(ctx *gin.Context) {
	var inquiry models.InquiryRequest

	if err := ctx.ShouldBindJSON(&inquiry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}



	
	
	err := models.CreateInquiry(api.db, inquiry)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	inquiryReq := ContactUs{
		FirstName:    inquiry.FirstName,
		LastName:     inquiry.LastName,
		Email:        inquiry.Email,
		PhoneNumber:  inquiry.Phone,
		Message:      inquiry.Message,
	}

	emailErr := api.SendContactUsMail(inquiryReq)
	
	if emailErr != nil {
		errorMessage := fmt.Sprintf("Error sending contact us email: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	// insert send email here to James@elated.io for prod / use your own email for dev

	ctx.JSON(http.StatusOK, gin.H{"message": "Inquiry created successfully"})
}

func (api *API) GetInquiries(ctx *gin.Context) {
	inquiries, err := models.GetAllInquiries(api.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inquiries)
}

func (api *API) GetInquiry(ctx *gin.Context) {
	id := ctx.Param("id")
	inquiry, err := models.GetInquiryByID(api.db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, inquiry)
}
