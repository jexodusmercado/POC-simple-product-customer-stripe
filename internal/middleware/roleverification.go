package middleware

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RequiresRole(tx *gorm.DB, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
