package api

import (
	"github.com/google/uuid"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) GetTransactions(c *gin.Context) {
	transactions, err := models.GetTransactions(api.db)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transactions)
}

func (api *API) GetTransaction(c *gin.Context) {
	id := c.Param("id")

	transactionID, err := uuid.Parse(id)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	transaction, err := models.GetTransactionByID(api.db, transactionID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}
