package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jexodusmercado/POC-simple-product-customer-stripe/internal/models"
)


func (api *API) GetQrCodeByUserIdAndTransactionId(ctx *gin.Context) {
	id := ctx.Param("id")
	transaction, err := models.GetTransactionByStripePaymentIntentID(api.db, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	qrcode, qrErr := models.GetQrCodeByUserIDAndTransactionID(api.db, transaction.ID, transaction.UserID)
	if qrErr != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, qrcode)
}
