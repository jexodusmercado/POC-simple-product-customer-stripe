package middleware

import (
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const (
	BEARER = "Bearer "
)

func RequiresAuthentication(tx *gorm.DB, env *conf.GlobalConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		

		c.Next()
	}
}
