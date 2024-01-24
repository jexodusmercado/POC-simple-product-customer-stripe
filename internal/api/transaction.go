package api

import (
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
	ID := c.Query("id")

	transaction, err := models.GetTransactionByID(api.db, ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, transaction)
}
