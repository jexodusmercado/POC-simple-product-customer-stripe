package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/conf"
)

func AddRequestID(env *conf.GlobalConfig) gin.HandlerFunc {
	return func(c *gin.Context) {
		id := ""
		if env.RequestIDHeader != "" {
			id = c.GetHeader(env.RequestIDHeader)
		}
		if id == "" {
			uid := uuid.New()
			id = uid.String()
		}

		c.Set("request_id", id)
	}
}
