package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)

func (api *API) CreateProduct(c *gin.Context) {

	var req models.CreateProductRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err := models.CreateProduct(api.db, &req)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, nil)
}

func (api *API) GetProducts(c *gin.Context) {
	products, err := models.GetProducts(api.db)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, products)
}

func (api *API) GetProduct(c *gin.Context) {
	ID := c.Query("id")

	productID, err := uuid.Parse(ID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	product, err := models.GetProductByID(api.db, productID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, product)
}
