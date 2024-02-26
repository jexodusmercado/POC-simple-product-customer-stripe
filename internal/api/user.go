package api

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := models.CreateUser(api.db, &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	userBetaReq := BetaList{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		Email:       req.Email,
		PhoneNumber: req.PhoneNumber,
		ZipCode:     req.ZipCode,
	}

	//send the beta email registration if isJoinBeta != nil
	if req.IsJoinBeta != nil {
		userBetaReq.IsJoinBeta = req.IsJoinBeta
		emailErr := api.SendBetaMail(userBetaReq)

		if emailErr != nil {
			errorMessage := fmt.Sprintf("Error sending contact us email: %v", err)
			c.JSON(http.StatusNotFound, gin.H{"error": errorMessage})
			return
		}
	}

	c.JSON(http.StatusOK, user)
}

func (api *API) GetUsers(c *gin.Context) {

	users, err := models.GetUsers(api.db)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, users)
}

func (api *API) GetUser(c *gin.Context) {
	ID := c.Query("id")

	userID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	user, err := models.GetUserByID(api.db, userID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}

func (api *API) CheckUserExists(c *gin.Context) {
	email := c.Param("email")

	exists, err := models.CheckUserExists(api.db, email)
	if err != nil {
		c.JSON(http.StatusOK, false)
		return
	}

	c.JSON(http.StatusOK, exists)
}


func (api *API) UnsubscribeUser(c *gin.Context) {
	email := c.Param("email")

	user, err := models.UnsubscribeUser(api.db, email)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, user)
}
