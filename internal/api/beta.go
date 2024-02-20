package api

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateBetaUser(ctx *gin.Context) {
	var betaUser models.CreateBetaUserRequest

	if err := ctx.ShouldBindJSON(&betaUser); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Check if the beta user already exists with the provided email
    existingBetaUser, err := models.GetBetaUserByEmail(api.db, betaUser.Email)

    if existingBetaUser != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Beta user with this email already exists"})
        return
    }

	err = models.CreateBetaUser(api.db, betaUser)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Assuming BetaList is the correct type for sending the email
	betaListReq := BetaList{
		FirstName:    betaUser.FirstName,
		LastName:     betaUser.LastName,
		Email:        betaUser.Email,
		PhoneNumber:  betaUser.Phone,
		ZipCode:      betaUser.Zipcode,
	}

	emailErr := api.SendBetaMail(betaListReq)

	if emailErr != nil {
		errorMessage := fmt.Sprintf("Error sending welcome email: %v", err)
		ctx.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Beta user created successfully"})
}

func (api *API) GetBetaUsers(ctx *gin.Context) {
	betaUsers, err := models.GetBetaUsers(api.db)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, betaUsers)
}

func (api *API) GetBetaUserById(ctx *gin.Context) {
	id := ctx.Param("id")
	betaUser, err := models.GetBetaUserById(api.db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, betaUser)
}

